package combinator

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/mcfly722/formulator/constants"
	"github.com/mcfly722/formulator/functions"
	"github.com/mcfly722/formulator/operators"
	"github.com/mcfly722/formulator/vm"
)

/*
func testRecombination(t *testing.T, testBracketSequence string) {
	t.Log(fmt.Sprintf("sequence: %v", testBracketSequence))

	i := 1
	readyProgram := func(program *vm.Program) {

		decompiled, err := vm.Decompile(program)
		if err != nil {
			t.Errorf("%v", err)
		}

		t.Log(fmt.Sprintf("%3v) %v          %v     %v     %v", i, constants.CombinationOfPointersToString(&program.Constants, " "), operators.CombinationToString(&program.Operators, " "), functions.CombinationToString(&program.Functions, " "), decompiled))
		i++

	}

	err := RecombineSequence(testBracketSequence, &constants.AvailableConstants, functions.Functions, operators.Operators, 1, 1, 1, false, readyProgram)
	if err != nil {
		t.Errorf("%v", err)
	}

}

func Test_Recombination1(t *testing.T) {
	testRecombination(t, "()")
}

func Test_Recombination2(t *testing.T) {
	testRecombination(t, "(())")
}

func Test_Recombination3(t *testing.T) {
	testRecombination(t, "()()")
}

func Test_Recombination4(t *testing.T) {
	testRecombination(t, "((()))")
}

func Test_Recombination5(t *testing.T) {
	testRecombination(t, "(()())")
}

func Test_Recombination6(t *testing.T) {
	testRecombination(t, "((()()))")
}

*/
const TotalPoints = 30

func Test_EXP(t *testing.T) {

	var sequence = "((()())((())()))()"

	var availableConstants = []float64{-1, constants.X, constants.N, constants.PreviousValue0, constants.PreviousValue1}
	var availableFunctions = []*functions.Function{&functions.Factorial}
	var availableOperators = []*operators.Operator{&operators.Add, &operators.Multiply, &operators.Power}

	points := []Point{}

	for i := 0; i < TotalPoints; i++ {
		x := rand.Float64() * 10
		point := Point{x, math.Exp(x)}
		points = append(points, point)

		t.Log(fmt.Sprintf("%2v) %2.8f %05.4f", i, point.X, point.Y))
	}

	t.Log(fmt.Sprintf("sequence: %v", sequence))

	i := 1

	var deviationThreshold float64 = 1000000
	var stopThreshold float64 = 1

	readyProgram := func(program *vm.Program) bool {

		deviation, err := CalculateDeviation(program, &points, deviationThreshold, 30, false)

		if err == nil {
			if deviation <= deviationThreshold {
				deviationThreshold = deviation

				decompiled, err := vm.Decompile(program)
				if err != nil {
					t.Errorf("%v", err)
				}
				fmt.Println(fmt.Sprintf("%3v) %v          %v     %v     %v", i, constants.CombinationOfPointersToString(&program.Constants, " "), operators.CombinationToString(&program.Operators, " "), functions.CombinationToString(&program.Functions, " "), decompiled))
				fmt.Println(fmt.Sprintf("deviation = %v", deviationThreshold))

			}
		}

		if deviation < stopThreshold == true {
			return false
		}

		i++
		return true
	}

	err := RecombineSequence(sequence, &availableConstants, availableFunctions, availableOperators, 2, 1, 1, true, readyProgram)
	if err != nil {
		t.Errorf("%v", err)
	}

}
