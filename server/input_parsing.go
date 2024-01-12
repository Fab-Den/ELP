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

type Inequalities struct {
	inequalities []Inequality
}

type Inequality struct {
	// left < right
	str   string
	left  Operation
	right Operation
}
type Operation struct {
	expression string
	operator   rune
	elements   []Operation
}

func (I *Inequality) evaluate(listVar []Variable, varValues []float64) bool {

	if I.left.getValue(listVar, varValues) < I.right.getValue(listVar, varValues) {
		return true
	} else {
		return false
	}
}

func (I *Inequalities) evaluate(listVar []Variable, varValues []float64) bool {
	for _, inequality := range I.inequalities {
		if inequality.evaluate(listVar, varValues) == false {
			return false
		}
	}
	return true
}

func (I *Inequality) innit() {
	I.left.expression = ""
	I.right.expression = ""

	i := 0
	for I.str[i] != '<' && I.str[i] != '>' {
		I.left.expression += string(I.str[i])
		i += 1
	}

	separator := I.str[i]

	i += 1

	for i < len(I.str) {
		I.right.expression += string(I.str[i])
		i += 1
	}

	if separator == '>' {
		I.left, I.right = I.right, I.left
	}

	I.left.innit()
	I.right.innit()
}

func (O *Operation) innit() {
	substring := O.expression

	ind := findIndexOfChar(substring, '(')
	ind2 := findLastIndexOfChar(substring, ')')

	if ind == 0 && ind2 == len(substring)-1 {
		O.elements = append(O.elements, Operation{expression: substring[1 : len(substring)-1]})
		O.operator = '('
	} else {
		op := 'n'
		indexOp := 0
		listOp := []rune{'-', '+', '*', '/'}

		for _, currentOp := range listOp {
			sub := O.expression

			currentOpIndex := findIndexOfChar(O.expression, currentOp)

			for currentOpIndex != -1 {

				if currentOpIndex != -1 && (currentOpIndex < ind || currentOpIndex > ind2) {
					op = currentOp
					indexOp = currentOpIndex
					break
				} else {
					sub = O.expression[currentOpIndex+1:]
					find := findIndexOfChar(sub, currentOp)
					if find == -1 {
						currentOpIndex = -1
					} else {
						currentOpIndex = currentOpIndex + 1 + find
					}

				}

			}
			if op != 'n' {
				break
			}
		}

		O.operator = op

		if op != 'n' {

			newO := Operation{expression: O.expression[0:indexOp]}
			if len(O.expression) != 0 {
				O.elements = append(O.elements, newO)
			}

			newO = Operation{expression: O.expression[indexOp+1:]}
			if len(O.expression) != 0 {
				O.elements = append(O.elements, newO)
			}

		}
	}
	if O.elements != nil && len(O.elements) > 0 {
		for i := range O.elements {
			O.elements[i].innit()
		}
	}
}

func (O *Operation) getValue(listVar []Variable, varValues []float64) float64 {

	if O.operator == 'n' {
		floatValue, err := strconv.ParseFloat(O.expression, 64)
		if err == nil {
			return floatValue
		}

		varIndex := findVariableIndex(listVar, O.expression)

		if varIndex != -1 {
			return varValues[varIndex]
		}

	} else {

		if O.operator == '*' {
			return O.elements[0].getValue(listVar, varValues) * O.elements[1].getValue(listVar, varValues)

		} else if O.operator == '/' {
			return O.elements[0].getValue(listVar, varValues) / O.elements[1].getValue(listVar, varValues)

		} else if O.operator == '+' {
			return O.elements[0].getValue(listVar, varValues) + O.elements[1].getValue(listVar, varValues)

		} else if O.operator == '-' {
			return O.elements[0].getValue(listVar, varValues) - O.elements[1].getValue(listVar, varValues)

		} else if O.operator == '(' {
			return O.elements[0].getValue(listVar, varValues)
		}
	}
	return 0
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

func initializeInequalities(file *os.File) (Inequalities, error) {
	scanner := bufio.NewScanner(file)
	Is := Inequalities{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[0] == '#' {
			if countCharInString(line, '<')+countCharInString(line, '>') == 1 {
				Is.inequalities = append(Is.inequalities, Inequality{str: line[1:]})
			}
		}
	}

	for i := range Is.inequalities {
		Is.inequalities[i].innit()
	}

	return Is, nil
}

func findVariableIndex(listVar []Variable, name string) int {
	for i, V := range listVar {
		if V.name == name {
			return i
		}
	}
	return -1
}

func findIndexOfChar(str string, charR rune) int {
	for i, char := range str {
		if char == charR {
			return i
		}
	}
	return -1
}

func findLastIndexOfChar(str string, charR rune) int {
	lastIndex := -1

	for i, char := range str {
		if char == charR {
			lastIndex = i
		}
	}
	return lastIndex
}

func countCharInString(str string, char rune) int {
	out := 0
	for _, v := range str {
		if v == char {
			out += 1
		}
	}

	return out
}
