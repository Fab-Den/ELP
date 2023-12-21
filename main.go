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

	for _, V := range listVar {

		_, err = file.Seek(0, 0)
		if err != nil {

		}

		err := V.setVariableRange(file)
		if err != nil {
			println(err.Error())
		}
	}

}
