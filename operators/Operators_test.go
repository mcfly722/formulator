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
	o := newOperator(func(a float64, b float64) float64 { return a + b }, "+")

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
	o := newOperator(func(a float64, b float64) float64 { return math.Pow(a, b) }, "^")

	c := o.Calculate(a, b)

	expression := fmt.Sprintf("%v = %v", o.OperatorExpressionToString(fmt.Sprintf("%f", a), fmt.Sprintf("%f", b)), c)

	if c != math.Pow(a, b) {
		t.Errorf(expression)
	} else {
		t.Log(expression)
	}

}

func Test_OperatorsRecombination(t *testing.T) {

	addOperator := newOperator(func(a float64, b float64) float64 { return a + b }, "+")
	multiplyOperator := newOperator(func(a float64, b float64) float64 { return a * b }, "*")
	powerOperator := newOperator(func(a float64, b float64) float64 { return math.Pow(a, b) }, "^")

	availableOperatorsTypes := []*Operator{addOperator, multiplyOperator, powerOperator}

	combination := []*Operator{nil, nil, nil, nil, nil}

	i := 1

	ready := func() {
		operators := []string{}
		for _, operator := range combination {
			operators = append(operators, operator.Separator)
		}
		t.Log(fmt.Sprintf("%4v) %v", i, strings.Join(operators, " ")))

		i++
	}

	Recombination(availableOperatorsTypes, combination, ready)
}
