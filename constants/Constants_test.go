package constants

import (
	"fmt"
	"math"
	"testing"
)

// go test -v constants.go constants_test.go

func Test_ConstantsCombinations(t *testing.T) {

	constants := []float64{2, 3, 5, -1, math.Pi}

	ready := func(constantsCombination *[]float64) {
		t.Log(fmt.Sprintf("%v", constantsCombination))
	}

	Recombination(&constants, ready)
}
