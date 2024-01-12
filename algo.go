package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// Generates a point with random coordinates
// all coordinates are kept in the range defined in the input file
func generatePoint(listVar []Variable, input_chan chan []float64, N int) {
	for j := 0; j < N; j++ {
		var point []float64
		for i := 0; i < len(listVar); i++ {
			vari := listVar[i]
			point = append(point, randomFloat(vari.ran[0], vari.ran[1]))
		}
		input_chan <- point
	}

	close(input_chan)

}

func parallelisation(listVar []Variable, N int, nb_goroutine int, I Inequalities) {
	var wg sync.WaitGroup

	input_chan := make(chan []float64, 10)
	res_chan := make(chan bool, 10)
	doneChannel := make(chan bool, 1)

	go generatePoint(listVar, input_chan, N)
	go recoverData(res_chan, doneChannel)

	for i := 0; i <= nb_goroutine; i++ {
		wg.Add(1)
		go worker(input_chan, res_chan, listVar, I, &wg)
	}

	wg.Wait()

	close(res_chan)

	<-doneChannel
}

func worker(input_chan <-chan []float64, res_chan chan<- bool, listVar []Variable, I Inequalities, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// get entry value
		point, more := <-input_chan
		if !more {
			// if channel close then quit
			return
		}
		res_chan <- I.evaluate(listVar, point)
	}
}

func recoverData(res_chan chan bool, doneChannel chan<- bool) {

	res := float64(0)
	c := 0

	for {
		data, more := <-res_chan
		if !more {
			fmt.Println(float64(res))

			doneChannel <- true
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
