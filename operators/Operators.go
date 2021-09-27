package operators

import (
	"fmt"
	"math"
)

// Add = a + b
const Add = 9237812 + 1

// Multiply = a * b
const Multiply = 9237812 + 2

// AbsPower = Abs(a) * b
const AbsPower = 9237812 + 3

// Operator structure represent operator
type Operator struct {
	oType int
}

func newOperator(operatorType int) *Operator {
	return &Operator{oType: operatorType}
}

// Calculate function return calculated expression with a and b
func (o Operator) Calculate(a float64, b float64) float64 {
	switch o.oType {
	case Add:
		return a + b
	case Multiply:
		return a * b
	case AbsPower:
		return math.Pow(math.Abs(a), b)
	default:
		return ^0
	}
}

// OperatorExpressionToString prints operator expression
func (o Operator) OperatorExpressionToString(a string, b string) string {
	switch o.oType {
	case Add:
		return fmt.Sprintf("%v + %v", a, b)
	case Multiply:
		return fmt.Sprintf("%v * %v", a, b)
	case AbsPower:
		return fmt.Sprintf("%v ^ %v", a, b)
	default:
		return fmt.Sprintf("unknown operator type = %v", o.oType)
	}
}
