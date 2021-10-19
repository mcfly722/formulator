package functions

import (
	"fmt"
	"math"
)

// Function structure
type Function struct {
	Name     string
	Function func(a float64) float64
}

// Functions all known functions
var Functions = []*Function{
	{
		Name:     "round",
		Function: func(a float64) float64 { return math.Round(a) },
	},
	{
		Name: "odd",
		Function: func(a float64) float64 {
			if int64(a)%2 == 1 {
				return 1
			}
			if a == 0 {
				return 0
			}
			return -1
		},
	},
	{
		Name:     "abs",
		Function: func(a float64) float64 { return math.Abs(a) },
	},
}

// FunctionByName get function by its name
func FunctionByName(name string) (*Function, error) {
	for n := range Functions {
		if Functions[n].Name == name {
			return Functions[n], nil
		}
	}
	return nil, fmt.Errorf(fmt.Sprintf("function %v is unknown", name))
}

// Calculate function
func (function Function) Calculate(argument float64) float64 {
	return function.Function(argument)
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
