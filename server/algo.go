package main

import (
	"math/rand"
)

type mainInputContainer struct {
	listVar       []Variable
	inequalities  Inequalities
	outputChannel chan subOutputContainer
	N             int
}

type subOutputContainer struct {
	N     int
	value int
}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func worker(mainInputChannel <-chan mainInputContainer) {
	for {
		element := <-mainInputChannel
		total := 0
		for i := 0; i < element.N; i++ {

			// Generates a point with random coordinates
			// all coordinates are kept in the range defined in the input file
			var point []float64
			for i := 0; i < len(element.listVar); i++ {
				variable := element.listVar[i]
				point = append(point, randomFloat(variable.ran[0], variable.ran[1]))
			}

			if element.inequalities.evaluate(element.listVar, point) == true {
				total += 1
			}
		}
		element.outputChannel <- subOutputContainer{N: element.N, value: total}
	}
}

func receiveDataForRequest(outputChannel <-chan subOutputContainer, resultChannel chan<- float64, N int) {
	totalN := 0
	totalValue := 0
	for totalN < N {
		temp := <-outputChannel
		totalN = totalN + temp.N
		totalValue = totalValue + temp.value
	}

	resultChannel <- float64(totalValue) / float64(totalN)

}
