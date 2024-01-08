package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Variable struct {
	name string
	ran  [2]float64
}

func (V *Variable) setVariableRange(file *os.File) error {

	scanner := bufio.NewScanner(file)
	rangeLine := ""

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > len(V.name)+3 {
			if line[0:len(V.name)+1] == "~"+V.name {
				rangeLine = line[len(V.name)+3:]
			}
		}
	}
	if rangeLine == "" {
		return fmt.Errorf("no range defined for %s variable", V.name)
	}
	splitRangeLine := strings.Split(rangeLine, ":")

	if len(splitRangeLine) != 2 {
		return fmt.Errorf("range of variable %s bad format", V.name)
	}

	for i := 0; i < 2; i++ {
		floatValue, err := strconv.ParseFloat(splitRangeLine[i], 64)

		if err != nil {
			return err
		}

		V.ran[i] = floatValue
	}

	if V.ran[0] > V.ran[1] {
		V.ran[0], V.ran[1] = V.ran[1], V.ran[0]
	}

	return nil
}

func initializeVariables(file *os.File) ([]Variable, error) {

	var listOfVariables []Variable

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		text := scanner.Text()
		names := strings.Split(text, " ")

		for _, value := range names {
			listOfVariables = append(listOfVariables, Variable{name: value})
		}

		return listOfVariables, nil

	} else if err := scanner.Err(); err != nil {
		return listOfVariables, err
	} else {
		return listOfVariables, fmt.Errorf("empty file")
	}
}
