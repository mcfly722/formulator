package constants

import (
	"fmt"
	"testing"
)

// go test -v constants.go constants_test.go

func Test_ConstantsCombinations(t *testing.T) {
	i := 1

	availableConstants := []float64{3, IterationIndex, PreviousIterationValue, Argument}

	combination := []float64{0, 0, 0}

	combinationPointer := []*float64{&combination[0], &combination[1], &combination[2]}

	ready := func(constantsCombination *[]*float64) {
		t.Log(fmt.Sprintf("%3v) %v", i, CombinationToString(&combination, " ")))
		i++
	}

	Recombination(&availableConstants, &combinationPointer, 1, 2, 3, true, ready)
}
