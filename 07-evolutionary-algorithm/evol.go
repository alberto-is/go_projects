package main

import (
	"math/rand"
	"strings"
	"time"
)

const target string = "Hola mundo"
const targetLength int = len(target)

const populationSize int = 100

const mutacion float32 = 0.01

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
	var bestCandidate1 Individual
	bestFitness1 := -1

	for i := 0; i < 5; i++ {
		candidate := (*population)[random.Intn(populationSize)]
		if candidate.fitness > bestFitness1 {
			bestCandidate1 = candidate
			bestFitness1 = candidate.fitness
		}
	}

	var bestCandidate2 Individual
	bestFitness2 := -1
	for i := 0; i < 5; i++ {
		candidate := (*population)[random.Intn(populationSize)]
		if candidate.fitness > bestFitness2 {
			bestCandidate2 = candidate
			bestFitness2 = candidate.fitness
		}
	}

	return [2]Individual{bestCandidate1, bestCandidate2}
}

func crossover() {

}

func mutate() {

}

func updatePopulation() {

}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	createInitialPopulation(random)
}
