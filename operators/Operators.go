package operators

import (
	"fmt"
)

// Operator structure
type Operator struct {
	f         func(a float64, b float64) float64
	separator string
}

// Constructor
func newOperator(calculator func(a float64, b float64) float64, _separator string) *Operator {
	return &Operator{
		f:         calculator,
		separator: _separator,
	}
}

// Calculate function
func (o Operator) Calculate(a float64, b float64) float64 {
	return o.f(a, b)
}

// OperatorExpressionToString string representation of operator
func (o Operator) OperatorExpressionToString(a string, b string) string {
	return fmt.Sprintf("%v %v %v", a, o.separator, b)
}
