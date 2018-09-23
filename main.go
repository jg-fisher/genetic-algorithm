package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type individual struct {
	genotype []int
	Fitness  int
}

func main() {
	var (
		population       []individual
		topPercent       int
		populationNum    int
		maxGenotypeLen   int
		generationsNum   int
		maxGenotypeValue int
	)

	// hyperparams
	topPercent = 40
	populationNum = 1000
	maxGenotypeLen = 10
	generationsNum = 100
	maxGenotypeValue = 100

	population = makePopulation(populationNum, maxGenotypeLen, maxGenotypeValue)
	population = eval_fitness(population)

	fmt.Println("Initial Population Winning fitness: ", population[0].Fitness)
	fmt.Println("Initial Population Winning genotype: ", population[0].genotype)

	for generation := 0; generation < generationsNum; generation++ {
		population = eval_fitness(population)
		population = selection(population, topPercent)
		population = crossover(population, populationNum)
		population = mutation(population, maxGenotypeValue)
	}

	population = selection(population, topPercent)
	fmt.Println("Final Population Winning fitness: ", population[0].Fitness)
	fmt.Println("Final Population Winning genotype: ", population[0].genotype)
}

func mutation(population []individual, maxGenotypeValue int) []individual {
	var (
		mutationRate float64
		randomNumber float64
	)

	mutationRate = 0.05
	for _, i := range population {
		randomNumber = rand.Float64()
		if mutationRate > randomNumber {
			i.genotype[rand.Intn(len(i.genotype)-1)] = rand.Intn(maxGenotypeValue)
		}
	}

	return population
}

func selection(population []individual, topPercent int) []individual {
	topPercent = len(population) * topPercent / 100

	// sort population by fitness
	sort.Slice(population[:], func(i, j int) bool {
		return population[i].Fitness > population[j].Fitness
	})

	population = population[0:topPercent]
	return population
}

// sum of genotype
func eval_fitness(population []individual) []individual {
	var sum int
	for i := 0; i < len(population); i++ {
		for gene := 0; gene < len(population[i].genotype); gene++ {
			sum += population[i].genotype[gene]
		}
		population[i].Fitness = sum
		sum = 0
	}
	return population
}

func crossover(population []individual, populationNum int) []individual {

	for i := 0; i < (populationNum - len(population) + 20); i++ {

		var (
			parentOne       individual
			parentTwo       individual
			populationSlice []individual
			parentOneIndex  int
		)

		// choose random parents
		parentOneIndex = rand.Intn(len(population) - 1)
		parentOne = population[parentOneIndex]
		populationSlice = remove(population, parentOneIndex)
		parentTwo = populationSlice[rand.Intn(len(populationSlice)-1)]

		population = append(population, makeCrossoverIndividual(parentOne, parentTwo, len(parentOne.genotype)))

	}

	return population
}

func remove(slice []individual, s int) []individual {
	return append(slice[:s], slice[s+1:]...)
}

func makeCrossoverIndividual(parentOne individual, parentTwo individual, maxGenotypeLen int) individual {
	var (
		parentOneGene  int
		parentTwoGene  int
		parentGenePool []int
		randomGene     int
		newIndividual  individual
		genotype       []int
		fitness        int
	)

	// uniform crossover
	for index := 0; index < maxGenotypeLen; index++ {

		parentOneGene = parentOne.genotype[index]
		parentTwoGene = parentTwo.genotype[index]
		parentGenePool = append(parentGenePool, parentOneGene, parentTwoGene)

		randomGene = parentGenePool[rand.Intn(1)]
		parentGenePool = parentGenePool[:0]

		genotype = append(genotype, randomGene)
	}

	fitness = 0

	newIndividual = individual{
		genotype: genotype,
		Fitness:  fitness,
	}

	return newIndividual
}

func makePopulation(populationNum int, maxGenotypeLen int, maxGenotypeValue int) []individual {
	var (
		population []individual
	)
	for i := 0; i < populationNum; i++ {
		population = append(population, makeRandomIndividual(maxGenotypeLen, maxGenotypeValue))
	}

	// can a regressor model be implemented to learn to generate crossovers?

	return population
}

func makeRandomIndividual(maxGenotypeLen int, maxGenotypeValue int) individual {
	var (
		newIndividual individual
		genotype      []int
		fitness       int
	)

	genotype = makeGenotype(maxGenotypeLen, maxGenotypeValue)
	fitness = 0
	newIndividual = individual{
		genotype: genotype,
		Fitness:  fitness,
	}

	return newIndividual
}

func makeGenotype(maxGenotypeLength int, maxGenotypeValue int) []int {
	var (
		genotype     []int
		randomNumber int
	)

	for gene := 0; gene < maxGenotypeLength; gene++ {
		randomNumber = rand.Intn(maxGenotypeValue)
		genotype = append(genotype, randomNumber)
	}

	return genotype
}
