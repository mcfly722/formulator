package operators

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"testing"
	"time"
)

// go test -v operators.go operators_test.go

func Test_Add(t *testing.T) {
	a := rand.Float64()*100 - 50
	b := rand.Float64()*100 - 50
	o, err := OperatorByName("+")
	if err != nil {
		t.Errorf(fmt.Sprintf("%v", err))
	}

	c := o.Calculate(a, b)

	expression := fmt.Sprintf("%v = %v", o.OperatorExpressionToString(fmt.Sprintf("%f", a), fmt.Sprintf("%f", b)), c)
	if c != a+b {
		t.Errorf(expression)
	} else {
		t.Log(expression)
	}

}

func Test_Power(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	a := rand.Float64() * 10 // positive only numbers :)))
	b := rand.Float64()*4 - 2
	o, err := OperatorByName("^")
	if err != nil {
		t.Errorf(fmt.Sprintf("%v", err))
	}

	c := o.Calculate(a, b)

	expression := fmt.Sprintf("%v = %v", o.OperatorExpressionToString(fmt.Sprintf("%f", a), fmt.Sprintf("%f", b)), c)

	if c != math.Pow(a, b) {
		t.Errorf(expression)
	} else {
		t.Log(expression)
	}

}

func Test_OperatorsRecombination(t *testing.T) {
	combination := []*Operator{nil, nil, nil, nil, nil}

	i := 1

	ready := func() {
		operators := []string{}
		for _, operator := range combination {
			operators = append(operators, operator.Name)
		}
		t.Log(fmt.Sprintf("%4v) %v", i, strings.Join(operators, " ")))

		i++
	}

	Recombination(Operators, combination, ready)
}
