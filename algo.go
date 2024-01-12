package main

import (
	"math/rand"
	"sync"
)

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func parallelization(listVar []Variable, N int, nbGoroutine int, I Inequalities) float64 {
	var wgWorkers sync.WaitGroup

	resultChannel := make(chan bool, 10)
	doneChannel := make(chan float64, 1)

	go recoverData(resultChannel, doneChannel)

	for i := 0; i <= nbGoroutine; i++ {
		wgWorkers.Add(1)
		go worker(N/nbGoroutine, resultChannel, listVar, I, &wgWorkers)
	}

	wgWorkers.Wait()
	close(resultChannel)

	return <-doneChannel
}

func worker(N int, resultChannel chan<- bool, listVar []Variable, I Inequalities, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < N; i++ {

		// Generates a point with random coordinates
		// all coordinates are kept in the range defined in the input file
		var point []float64
		for i := 0; i < len(listVar); i++ {
			variable := listVar[i]
			point = append(point, randomFloat(variable.ran[0], variable.ran[1]))
		}

		// Evaluate if the point is in the volume and put the result in the channel
		resultChannel <- I.evaluate(listVar, point)

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
