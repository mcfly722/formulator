package operators

import (
	"fmt"
)

// Operator structure
type Operator struct {
	f         func(a float64, b float64) float64
	Separator string
}

// Constructor
func newOperator(calculator func(a float64, b float64) float64, _separator string) *Operator {
	return &Operator{
		f:         calculator,
		Separator: _separator,
	}
}

// Calculate function
func (o Operator) Calculate(a float64, b float64) float64 {
	return o.f(a, b)
}

// OperatorExpressionToString string representation of operator
func (o Operator) OperatorExpressionToString(a string, b string) string {
	return fmt.Sprintf("%v %v %v", a, o.Separator, b)
}

func recombine(combination []*Operator, availableOperators []*Operator, depth int, ready func(current []*Operator)) {
	if depth > 0 {
		for _, operator := range availableOperators {
			newCombination := append(combination, operator)
			recombine(newCombination, availableOperators, depth-1, ready)
		}
	} else {
		ready(combination)
	}
}

// Recombination for operators
func Recombination(availableOperators []*Operator, combinationLenght int, ready func(current []*Operator)) {
	recombine([]*Operator{}, availableOperators, combinationLenght, ready)
}
