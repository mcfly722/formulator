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

	round, err := FunctionByName("round")
	if err != nil {
		t.Errorf(fmt.Sprintf("%v", err))
	}

	c := round.Calculate(argument)

	expression := fmt.Sprintf("%v = %v", round.FunctionExpressionToString(fmt.Sprintf("%f", argument)), c)

	if c != math.Round(argument) {
		t.Errorf(expression)
	} else {
		t.Log(expression)
	}

}

func Test_FunctionsRecombination(t *testing.T) {

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

	Recombination(Functions, combination, ready)
}
