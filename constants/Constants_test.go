package constants

import (
	"fmt"
	"testing"
)

// go test -v constants.go constants_test.go

func Test_ConstantsCombinations(t *testing.T) {
	i := 1

	constants := []float64{constantIterationIndex, constantPreviousIterationValue, constantArgument}

	ready := func(constantsCombination []float64) {
		t.Log(fmt.Sprintf("%5v %v", i, constantsCombination))
		i++
	}

	Recombination(&constants, 3, 1, 2, 3, ready)
}
