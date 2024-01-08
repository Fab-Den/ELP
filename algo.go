package main

import (
	"fmt"
	"math/rand"
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
}

func parallelisation(listVar []Variable, N int, nb_goroutine int) {
	input_chan := make(chan []float64, 10)
	res_chan := make(chan bool, 10)
	go generatePoint(listVar, input_chan, N)

	for i := 0; i <= nb_goroutine; i++ {
		go worker(input_chan, res_chan)
	}
	recoverData(res_chan, N)

}

func worker(input_chan chan []float64, res_chan chan bool) {
	point := <-input_chan
	res := false
	if point[0] < 0 {
		res = true
	}
	res_chan <- res
	go worker(input_chan, res_chan)
}

func recoverData(res_chan chan bool, N int) {
	res := 0
	for i := 0; i < N; i++ {
		if <-res_chan {
			res += 1
		}
	}
	fmt.Println(float64(res) / float64(N))
}
