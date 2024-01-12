package main

import (
	"fmt"
	"os"
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

	for _, V := range listVar {
		println(V.name, V.ran[0], V.ran[1])
	}

	_, err = file.Seek(0, 0)
	inequalities, err := initializeInequalities(file)

	println(inequalities.evaluate(listVar, []float64{0.45, 0.45}))
	parallelisation(listVar, 100000000, 10, inequalities)
}
