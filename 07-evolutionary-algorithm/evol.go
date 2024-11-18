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

	for _, indi := range population {
		indi = Individual{
			chros: randomString(random, targetLength),
		}
		indi.fitness = fitnessFunc(indi.chros)
	}

	return population
}

func arrangeForFitness() {

}
func seleccion() {

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
