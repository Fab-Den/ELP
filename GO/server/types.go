package main

// Variable is used to represent a variable from the formulas given in input. The structure is used to associate
// the variable to the range.
type Variable struct {
	name string
	ran  [2]float64
}

// Inequality represents one formula given in input. It stores the two members of the inequality. The left member is
// lower than the right member.
type Inequality struct {
	str   string
	left  Operation
	right Operation
}

// Operation stores an operation from the formula. The expression attribute is the string associated to the operation.
// Operator is a character to know which mathematical operation is associated. The 'n' value in this attribute
// represents the fact that there is not an operation to carry out, and the object represents then a single value.
// As an operation is carried out between two members, the elements attribute can store two other Operation structure.
type Operation struct {
	expression string
	operator   rune
	elements   []Operation
}

// Problem is a type that store all the information of a problem
type Problem struct {
	listVars       []Variable
	inequalities   []Inequality
	numberOfPoints int
}
