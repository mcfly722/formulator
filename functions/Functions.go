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

func recombine(combination []*Function, availableFunctions []*Function, depth int, ready func(current []*Function)) {
	if depth > 0 {
		for _, function := range availableFunctions {
			newCombination := append(combination, function)
			recombine(newCombination, availableFunctions, depth-1, ready)
		}
	} else {
		ready(combination)
	}
}

// Recombination for functions
func Recombination(availableFunctions []*Function, combinationLenght int, ready func(current []*Function)) {
	recombine([]*Function{}, availableFunctions, combinationLenght, ready)
}
