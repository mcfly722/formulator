package functions

import "fmt"

// Function structure
type Function struct {
	calculator func(a float64) float64
	name       string
}

// Constructor
func newFunction(_calculator func(a float64) float64, _name string) *Function {
	return &Function{
		calculator: _calculator,
		name:       _name,
	}
}

// Calculate function
func (function Function) Calculate(argument float64) float64 {
	return function.calculator(argument)
}

// FunctionExpressionToString string representation of function
func (function Function) FunctionExpressionToString(argument string) string {
	return fmt.Sprintf("%v(%v)", function.name, argument)
}
