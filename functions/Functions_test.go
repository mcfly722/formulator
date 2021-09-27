package functions

import (
	"fmt"
	"math"
	"math/rand"
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
