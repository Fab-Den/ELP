package main

import (
	"fmt"
	"strconv"
	"strings"
)

// evaluate is a method for the Inequality structure, which is called when we want to know if the inequality is
// respected according to the values of variables given in parameters.
func (I *Inequality) evaluate(listVar []Variable, varValues []float64) bool {

	if I.left.getValue(listVar, varValues) < I.right.getValue(listVar, varValues) {
		return true
	} else {
		return false
	}
}

// evaluate is a method for Problem that calls evaluate for each Inequality. Returns if all the inequalities
// are respected.
func (P *Problem) evaluate(listVar []Variable, varValues []float64) bool {
	for _, inequality := range P.inequalities {
		if !inequality.evaluate(listVar, varValues) {
			return false
		}
	}
	return true
}

// init is a method for Inequality, which parse the inequality string into the Operation instances stored into the left
// and right attributes. At the end of the parsing, call the init for each Operation instance.
func (I *Inequality) init() {
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

	I.left.init()
	I.right.init()
}

// init is a method for Operation which parse the expression string attribute of the object into the element attribute.
func (O *Operation) init() {
	substring := O.expression

	// we get the index of the first opening parenthesis and the index of the last closing parenthesis
	ind := findIndexOfChar(substring, '(')
	ind2 := findLastIndexOfChar(substring, ')')

	// if parenthesis encompass all the formula, we create a single Operation with the content of parenthesis as
	// expression attribute
	if ind == 0 && ind2 == len(substring)-1 {
		O.elements = append(O.elements, Operation{expression: substring[1 : len(substring)-1]})
		O.operator = '('
	} else {
		op := 'n'
		indexOp := 0
		listOp := []rune{'-', '+', '*', '/'}

		// in the order of listOp, we get the first operation outside parenthesis
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

		// if there is an operation to carry out, we add the left and right member as Operation instance in the
		// elements attribute
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

	// if there are Operation instances in element attributes, we initialize all of them
	if O.elements != nil && len(O.elements) > 0 {
		for i := range O.elements {
			O.elements[i].init()
		}
	}
}

// getValue is a method for Operation which gets the value of an Operation object according to values for variables,
// that are given in parameters. Returns then the value.
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

// setVariableRange is a method vor Variable which sets the ran attribute of Variable object according to the input
// string.
func (V *Variable) setVariableRange(str string) error {

	lines := strings.Split(str, "\n")

	rangeLine := ""

	for _, line := range lines {

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

// initializeVariables returns a list of variables according to the input string.
func initializeVariables(str string) ([]Variable, error) {

	var listOfVariables []Variable

	lines := strings.Split(str, "\n")

	names := strings.Split(lines[0], " ")

	for _, value := range names {
		listOfVariables = append(listOfVariables, Variable{name: value})
	}

	return listOfVariables, nil
}

// initInequalities is a method for Problem that edit the inequalities list according to the input string.
// Also, it calls the init method for each Inequality instance.
func (P *Problem) initInequalities(str string) error {
	lines := strings.Split(str, "\n")

	for _, line := range lines {
		if len(line) > 0 && line[0] == '#' {
			if countCharInString(line, '<')+countCharInString(line, '>') == 1 {
				P.inequalities = append(P.inequalities, Inequality{str: line[1:]})
			} else {
				return fmt.Errorf("bad number of inequality symbol into at least one formula")
			}
		}
	}

	for i := range P.inequalities {
		P.inequalities[i].init()
	}

	return nil
}

// getNumberOfPoints is a function that gets the number of points requested in the input file and returns the value.
func getNumberOfPoints(str string) (int, error) {
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		if strings.Contains(line, "N=") {
			value, err := strconv.Atoi(line[2:])
			if err != nil {
				return 0, fmt.Errorf("error when reading number of points : %s", err)
			}
			return value, nil
		}
	}
	return 0, fmt.Errorf("error when reading number of points : could not find number of points in data")
}

// getSpaceVolume is a function that returns the total volume of the space, according to the range of each variable.
func getSpaceVolume(listVars []Variable) float64 {
	total := float64(1)

	for _, variable := range listVars {
		total *= variable.ran[1] - variable.ran[0]
	}

	return total

}

// initVariableRange is a function that calls the setVariableRange for each Variable in listVar.
func initVariableRange(listVar []Variable, str string) error {

	for i := range listVar {
		err := listVar[i].setVariableRange(str)

		if err != nil {
			return err
		}
	}
	return nil
}

// findVariableIndex is a function that returns the index of a variable in a list of variables according to its name.
// returns -1 if the variable is not found.
func findVariableIndex(listVar []Variable, name string) int {
	for i, V := range listVar {
		if V.name == name {
			return i
		}
	}
	return -1
}

// findIndexOfChar is a function that returns the first index of the rune charR in the string str, or -1 if str doesn't
// contain the charR.
func findIndexOfChar(str string, charR rune) int {
	for i, char := range str {
		if char == charR {
			return i
		}
	}
	return -1
}

// findLastIndexOfChar is a function that returns the last index of the rune charR in the string str, or -1 if str
// doesn't contain the charR.
func findLastIndexOfChar(str string, charR rune) int {
	lastIndex := -1

	for i, char := range str {
		if char == charR {
			lastIndex = i
		}
	}
	return lastIndex
}

// countCharInString is a function that counts the number of rune char that the string str contains.
func countCharInString(str string, char rune) int {
	out := 0
	for _, v := range str {
		if v == char {
			out += 1
		}
	}

	return out
}

// getOperationSize is a method for Operation that returns the number of operations that the object contains. A single
// value count as 1 operation.
func (O *Operation) getOperationSize() int {
	total := 1
	for _, o := range O.elements {
		total += o.getOperationSize()
	}
	return total
}

// getProblemSize is a method for Problem that calculates the total number of operations among all inequalities.
func (P *Problem) getProblemSize() int {
	total := 0
	for _, i := range P.inequalities {
		total += i.left.getOperationSize() + i.right.getOperationSize() + 1
	}
	return total
}

func (O *Operation) checkData(listVar []Variable) error {
	if O.operator == 'n' {
		inVars := false

		for _, varName := range listVar {
			if varName.name == O.expression {
				inVars = true
				break
			}
		}

		if inVars == false {
			_, err := strconv.ParseFloat(O.expression, 64)

			if err != nil {
				return fmt.Errorf("bad char in formula : %s", O.expression)
			}
		}

	}
	return nil
}

func (P *Problem) checkData() error {
	var err error

	for _, i := range P.inequalities {

		err = i.left.checkData(P.listVars)
		if err != nil {
			return err
		}

		err = i.right.checkData(P.listVars)
		if err != nil {
			return err
		}

	}

	return nil
}
