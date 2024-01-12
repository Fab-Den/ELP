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

//func parallelization(listVar []Variable, N int, nbGoroutine int, I Inequalities) float64 {
//	var wgWorkers sync.WaitGroup
//
//	resultChannel := make(chan bool, 10)
//	doneChannel := make(chan float64, 1)
//
//	go recoverData(resultChannel, doneChannel)
//
//	for i := 0; i <= nbGoroutine; i++ {
//		wgWorkers.Add(1)
//		go worker(N/nbGoroutine, resultChannel, listVar, I, &wgWorkers)
//	}
//
//	wgWorkers.Wait()
//	close(resultChannel)
//
//	return <-doneChannel
//}

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

func recoverData(resultChannel chan bool, doneChannel chan<- float64) {

	res := float64(0)
	c := 0

	for {
		data, more := <-resultChannel

		if !more {
			doneChannel <- res
			return
		}

		if data {
			res = (res*float64(c) + 1) / (float64(c) + 1)
		} else {
			res = (res * float64(c)) / (float64(c) + 1)
		}
		c += 1

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
