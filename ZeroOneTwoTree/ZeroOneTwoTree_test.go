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

	Recombine(4, 2, ready)
}

func testBracketsForError(t *testing.T, brackets string) {
	_, err := GetNextTree(brackets)
	if err != nil {
		t.Log(fmt.Sprintf("correct error handling for %v -> %v", brackets, err))
	} else {
		t.Errorf("GetNextTree('%v') not returned error", brackets)
	}
}

func Test_BracketsOpensCloses(t *testing.T) {
	testBracketsForError(t, "(())(")
}

func Test_BracketsUnexpectedSymbol1(t *testing.T) {
	testBracketsForError(t, "(())!()")
}

func Test_BracketsUnexpectedSymbol2(t *testing.T) {
	testBracketsForError(t, "(())(!)")
}

func Test_BracketsUnexpectedSymbol3(t *testing.T) {
	testBracketsForError(t, "(())()!")
}

func Test_BracketsUnexpectedSymbol4(t *testing.T) {
	testBracketsForError(t, "!(())()")
}

func Test_BracketsClosesGreaterThanOpens(t *testing.T) {
	testBracketsForError(t, "((())))()")
}

func Test_GetNextTree(t *testing.T) {

	tree := "(()())()"

	nextTree, err := GetNextTree(tree)
	if err != nil {
		t.Errorf("GetNextTree('%v') returned error:%v", tree, err)
	}

	t.Log(fmt.Sprintf("%v -> %v", tree, nextTree))

}
