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

func recombine(availableOperators []*Operator, combination []*Operator, position int, ready func()) {
	if position < len(combination) {
		for _, operator := range availableOperators {
			combination[position] = operator
			recombine(availableOperators, combination, position+1, ready)
		}
	} else {
		ready()
	}
}

// Recombination for operators
func Recombination(availableOperators []*Operator, combination []*Operator, ready func()) {
	recombine(availableOperators, combination, 0, ready)
}
