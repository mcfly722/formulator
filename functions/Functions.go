package functions

import "fmt"

// Function structure
type Function struct {
	calculator func(a float64) float64
	Name       string
}

// Constructor
func newFunction(_calculator func(a float64) float64, _name string) *Function {
	return &Function{
		calculator: _calculator,
		Name:       _name,
	}
}

// Calculate function
func (function Function) Calculate(argument float64) float64 {
	return function.calculator(argument)
}

// FunctionExpressionToString string representation of function
func (function Function) FunctionExpressionToString(argument string) string {
	return fmt.Sprintf("%v(%v)", function.Name, argument)
}

func recombine(availableFunctions []*Function, combination []*Function, position int, ready func()) {
	if position < len(combination) {
		for _, function := range availableFunctions {
			(combination)[position] = function
			recombine(availableFunctions, combination, position+1, ready)
		}
	} else {
		ready()
	}
}

// Recombination for functions
func Recombination(availableFunctions []*Function, combination []*Function, ready func()) {
	recombine(availableFunctions, combination, 0, ready)
}
