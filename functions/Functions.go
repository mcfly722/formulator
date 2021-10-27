package functions

import (
	"fmt"
	"math"
	"strings"
)

// Function structure
type Function struct {
	Name     string
	Function func(a float64) float64
}

// Round function
var Round = Function{
	Name:     "round",
	Function: func(x float64) float64 { return math.Round(x) },
}

// Odd funciton
var Odd = Function{
	Name: "odd",
	Function: func(x float64) float64 {
		if int64(x)%2 == 1 {
			return 1
		}
		if x == 0 {
			return 0
		}
		return -1
	},
}

// Abs function
var Abs = Function{
	Name:     "abs",
	Function: func(x float64) float64 { return math.Abs(x) },
}

// Factorial function
var Factorial = Function{
	Name: "fact",
	Function: func(x float64) float64 {
		var result float64 = 1
		var i float64

		for i = 1; i <= x; i++ {
			result = result * i
			if math.IsNaN(result) || math.IsInf(result, 1) || math.IsInf(result, -1) {
				return result
			}
		}
		return result
	},
}

// Functions all known functions
var Functions = []*Function{&Round, &Odd, &Abs, &Factorial}

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

func recombine(availableFunctions []*Function, combination []*Function, position int, ready func(functionsCombination []*Function)) {
	if position < len(combination) {
		for _, function := range availableFunctions {
			(combination)[position] = function
			recombine(availableFunctions, combination, position+1, ready)
		}
	} else {
		ready(combination)
	}
}

// Recombination for functions
func Recombination(availableFunctions []*Function, combination []*Function, ready func(functionsCombination []*Function)) {
	recombine(availableFunctions, combination, 0, ready)
}

// CombinationToString string representation
func CombinationToString(combination *[]*Function, separator string) string {
	out := []string{}
	for _, function := range *combination {
		out = append(out, fmt.Sprintf("%5v", function.Name))
	}

	return strings.Join(out, separator)
}
