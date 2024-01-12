package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	file, err := os.Open("input_file.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	_, err = file.Seek(0, 0)
	if err != nil {

	}
	listVar, err := initializeVariables(file)

	for i := range listVar {

		_, err = file.Seek(0, 0)
		if err != nil {

		}

		err := listVar[i].setVariableRange(file)

		if err != nil {
			println(err.Error())
		}
	}

	_, err = file.Seek(0, 0)
	inequalities, err := monte_carlo.initializeInequalities(file)

	N := 10000000

	for i := 1; i < 20; i++ {
		timeStart := time.Now()
		println(main2.parallelization(listVar, N, i, inequalities))
		println("Execution time for ", N, " points over ", i, " go routines : ", time.Since(timeStart)/time.Millisecond, "ms")
	}

}
