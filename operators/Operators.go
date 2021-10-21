package operators

import (
	"fmt"
	"math"
	"strings"
)

// Operator structure
type Operator struct {
	Function func(a float64, b float64) float64
	Name     string
}

// Operators all known operators
var Operators = []*Operator{
	{
		Name:     "+",
		Function: func(a float64, b float64) float64 { return a + b },
	},
	{
		Name:     "*",
		Function: func(a float64, b float64) float64 { return a * b },
	},
	{
		Name:     "^",
		Function: func(a float64, b float64) float64 { return math.Pow(a, b) },
	},
}

// OperatorByName get operator by its name
func OperatorByName(name string) (*Operator, error) {
	for n := range Operators {
		if Operators[n].Name == name {
			return Operators[n], nil
		}
	}
	return nil, fmt.Errorf(fmt.Sprintf("operator %v is unknown", name))
}

// Calculate function
func (o Operator) Calculate(a float64, b float64) float64 {
	return o.Function(a, b)
}

// OperatorExpressionToString string representation of operator
func (o Operator) OperatorExpressionToString(a string, b string) string {
	return fmt.Sprintf("%v %v %v", a, o.Name, b)
}

func recombine(availableOperators []*Operator, combination []*Operator, position int, ready func(combination []*Operator)) {
	if position < len(combination) {
		for _, operator := range availableOperators {
			combination[position] = operator
			recombine(availableOperators, combination, position+1, ready)
		}
	} else {
		ready(combination)
	}
}

// Recombination for operators
func Recombination(availableOperators []*Operator, combination []*Operator, ready func(combination []*Operator)) {
	recombine(availableOperators, combination, 0, ready)
}

// CombinationToString string representation
func CombinationToString(combination *[]*Operator, separator string) string {
	out := []string{}
	for _, operator := range *combination {
		out = append(out, operator.Name)
	}

	return strings.Join(out, separator)
}
