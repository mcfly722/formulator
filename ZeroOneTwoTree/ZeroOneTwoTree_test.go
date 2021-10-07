package zeroOneTwoTree

import (
	"fmt"
	"testing"
)

func Test_Recombines(t *testing.T) {

	i := 1

	ready := func(bracketsStack []Point, diagonal [32]int) {
		representation := ""
		for _, point := range bracketsStack {
			for opens := 0; opens < point.Opens; opens++ {
				representation += "("
			}
			for closes := 0; closes < point.Closes; closes++ {
				representation += ")"
			}
		}

		t.Log(fmt.Sprintf("%5v   %v   %v", i, representation, diagonal))
		i++
	}

	Recombine(6, ready)
}
