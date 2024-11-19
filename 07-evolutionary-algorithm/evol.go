package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const target string = "En un lugar de la Mancha de cuyo nombre no quiero acordarme no ha mucho tiempo que vivia un hidalgo de los de lanza en astillero adarga antigua rocin flaco y galgo corredor"

const targetLength int = len(target)

const populationSize int = 100

const mutacion float32 = 0.01

const generations int = 10000000

type Individual struct {
	chros   string
	fitness int
}

func randomString(random *rand.Rand, length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ")
	var builder strings.Builder

	for range length {
		number := random.Intn(len(letters))
		builder.WriteRune(letters[number])
	}

	return builder.String()

}

func fitnessFunc(str string) int {
	var count int = 0
	for i := range str {
		if str[i] == target[i] {
			count++
		}
	}
	return count
}

func createInitialPopulation(random *rand.Rand) []Individual {
	population := make([]Individual, populationSize)

	for i := range population {
		population[i] = Individual{
			chros: randomString(random, targetLength),
		}
		population[i].fitness = fitnessFunc(population[i].chros)
	}

	return population
}

func seleccion(random *rand.Rand, population *[]Individual) [2]Individual {
	candidateChan := make(chan Individual, 2)
	defer close(candidateChan)

	findBest := func() {
		bestFitness := -1
		var bestCandidate Individual
		for i := 0; i < 5; i++ {
			candidate := (*population)[random.Intn(populationSize)]
			if candidate.fitness > bestFitness {
				bestCandidate = candidate
				bestFitness = candidate.fitness
			}
		}
		candidateChan <- bestCandidate
	}

	go findBest()
	go findBest()

	return [2]Individual{<-candidateChan, <-candidateChan}
}

func crossover(random *rand.Rand, parent1, parent2 Individual) Individual {
	crossoverPoint := random.Intn(targetLength)
	child := Individual{
		chros: parent1.chros[:crossoverPoint] + parent2.chros[crossoverPoint:],
	}
	child.fitness = fitnessFunc(child.chros)
	return child
}

func mutate(random *rand.Rand, individual Individual) Individual {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ")
	chros := []rune(individual.chros)

	for i := range chros {
		if random.Float32() < mutacion {
			chros[i] = letters[random.Intn(len(letters))]
		}
	}

	mutated := Individual{
		chros: string(chros),
	}
	mutated.fitness = fitnessFunc(mutated.chros)
	return mutated
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	population := createInitialPopulation(random)

	for gen := 0; gen < generations; gen++ {
		newPopulation := make([]Individual, populationSize)

		for i := 0; i < populationSize; i++ {
			parents := seleccion(random, &population)
			child := crossover(random, parents[0], parents[1])
			child = mutate(random, child)
			newPopulation[i] = child
		}

		population = newPopulation

		bestIndividual := population[0]
		for _, ind := range population {
			if ind.fitness > bestIndividual.fitness {
				bestIndividual = ind
			}
		}

		if bestIndividual.chros == target {
			fmt.Printf("Solution found in generation %d: %s\n", gen, bestIndividual.chros)
			break
		}

		if gen%100 == 0 {
			fmt.Printf("Generation %d: Best fitness = %d, Best string = %s\n",
				gen, bestIndividual.fitness, bestIndividual.chros)
		}
	}
}
