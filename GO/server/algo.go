package main

import (
	"math/rand"
)

// Type passed through the mainInputChannel
type mainInputContainer struct {
	problem       Problem
	outputChannel chan subOutputContainer
	N             int
}

// Type passed through channel between workers and result receiver
type subOutputContainer struct {
	N     int
	value int
}

// randomFloat is a function that generate a random number between min and max
func randomFloat(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// worker is a function that resolves the problem piece get from mainInputChannel
func worker(mainInputChannel <-chan mainInputContainer) {
	for {
		element := <-mainInputChannel
		total := 0
		for i := 0; i < element.N; i++ {

			// Generates a point with random coordinates
			// all coordinates are kept in the range defined in the input file
			var point []float64
			for i := 0; i < len(element.problem.listVars); i++ {
				variable := element.problem.listVars[i]
				point = append(point, randomFloat(variable.ran[0], variable.ran[1]))
			}

			// Evaluate the point and increase the total if true
			if element.problem.evaluate(element.problem.listVars, point) {
				total += 1
			}
		}
		element.outputChannel <- subOutputContainer{N: element.N, value: total}
	}
}

// receiveDataForRequest is a function that gets the results of workers from outputChannel, gathers them and send a
// global result in a single message through resultChannel.
func receiveDataForRequest(outputChannel <-chan subOutputContainer, resultChannel chan<- float64, N int) {
	// N is the number of point that have to be evaluated
	// totalN is the number of points for which we received a result
	// totalValue is the number of points that are evaluated as true

	totalN := 0
	totalValue := 0
	for totalN < N {
		temp := <-outputChannel
		totalN = totalN + temp.N
		totalValue = totalValue + temp.value
	}

	// send the proportion of points that fulfill the criteria
	resultChannel <- float64(totalValue) / float64(totalN)

}
