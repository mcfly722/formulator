package constants

import (
	"fmt"
	"testing"
)

// go test -v constants.go constants_test.go

func Test_ConstantsCombinations(t *testing.T) {
	i := 1

	combination := []float64{0, 0, 0}
	combinationPointers := []*float64{&combination[0], &combination[1], &combination[2]}

	ready := func(constantsCombination *[]*float64) {
		t.Log(fmt.Sprintf("%3v) %v", i, CombinationToString(&combination, " ")))
		i++
	}

	Recombination(&AvailableConstants, &combinationPointers, 1, 2, 3, true, ready)
}
