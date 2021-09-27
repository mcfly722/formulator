package functions

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"testing"
)

func Test_FunctionsRound(t *testing.T) {
	argument := rand.Float64()*100 - 50

	f := newFunction(func(a float64) float64 {
		return math.Round(a)
	}, "round")

	c := f.Calculate(argument)

	expression := fmt.Sprintf("%v = %v", f.FunctionExpressionToString(fmt.Sprintf("%f", argument)), c)

	if c != math.Round(argument) {
		t.Errorf(expression)
	} else {
		t.Log(expression)
	}

}

func Test_FunctionsRecombination(t *testing.T) {

	roundFunction := newFunction(func(a float64) float64 { return math.Round(a) }, "round")

	// returns 0 = 0, if odd = 1, even = -1
	oddFunction := newFunction(func(a float64) float64 {
		if int64(a)%2 == 1 {
			return 1
		}
		if a == 0 {
			return 0
		}
		return -1
	}, "  odd")

	availableFunctionsTypes := []*Function{roundFunction, oddFunction}

	combination := []*Function{nil, nil, nil, nil, nil}

	i := 1
	ready := func() {

		functions := []string{}
		for _, function := range combination {
			functions = append(functions, function.Name)
		}
		t.Log(fmt.Sprintf("%4v) %v", i, strings.Join(functions, " , ")))

		i++
	}

	Recombination(availableFunctionsTypes, combination, ready)
}
