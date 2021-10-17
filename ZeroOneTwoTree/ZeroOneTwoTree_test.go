package zeroOneTwoTree

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Recombines(t *testing.T) {

	i := 1

	ready := func(bracketsStack []BracketStep, diagonal [32]int) {
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
	_, err := GetNextTree(brackets, 2)
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

func Test_BracketsToTree(t *testing.T) {
	bracketSequence := "()((())())"
	expression, err := BracketsToTree(bracketSequence)
	if err != nil {
		t.Errorf("Cant build expression tree for %v. Error: %v", bracketSequence, err)
	}

	bytes, err := json.Marshal(expression)
	if err != nil {
		t.Errorf("Can't serialize %v. Error:%v", bracketSequence, err)
	}
	t.Log(string(bytes))

}

func testBracketsToTreeError(t *testing.T, testSequence string) {
	_, err := BracketsToTree(testSequence)

	if err == nil {
		t.Errorf("error for '%v' sequence does not catched", testSequence)
	} else {
		t.Log(fmt.Sprintf("error for '%v' successfully catched.\nerror description:%v", testSequence, err))
	}
}

func testBracketsToTreeSuccess(t *testing.T, testSequence string) {
	expression, err := BracketsToTree(testSequence)

	if err != nil {
		t.Errorf("error for %v :%v", testSequence, err)
	}

	bytes, err := json.Marshal(expression)
	if err != nil {
		t.Errorf("Can't serialize %v. Error:%v", testSequence, err)
	}
	t.Log(string(bytes))
}

func Test_BracketsToTree_IncorrectSymbolFromStart(t *testing.T) {
	testBracketsToTreeError(t, "a(()(()))")
}

func Test_BracketsToTree_IncorrectSymbolAtTheEnd(t *testing.T) {
	testBracketsToTreeError(t, "(()(()))b")
}

func Test_BracketsToTree_IncorrectSymbolInTheMiddle(t *testing.T) {
	testBracketsToTreeError(t, "(()((c)))")
}

func Test_BracketsToTree_LostOpeningBracket(t *testing.T) {
	testBracketsToTreeError(t, "())")
}

func Test_BracketsToTree_LostClosingBracket(t *testing.T) {
	testBracketsToTreeError(t, "(()")
}

func Test_BracketsToTree_WrongBracketsSequence(t *testing.T) {
	testBracketsToTreeError(t, "())(")
}

func Test_BracketsToTree_FirstBracket(t *testing.T) {
	testBracketsToTreeSuccess(t, "()")
}

func Test_GetNextTree(t *testing.T) {
	bracketSequence := "()"

	for i := 1; i < 200; i++ {

		tree, err := BracketsToTree(bracketSequence)
		if err != nil {
			t.Errorf("GetNextTree('%v') returned error:%v", bracketSequence, err)
		}

		max := tree.MaxChilds()

		nextBracketSequcence, err := GetNextTree(bracketSequence, 1000)
		if err != nil {
			t.Errorf("GetNextTree('%v') returned error:%v", bracketSequence, err)
		}

		t.Log(fmt.Sprintf("%3v) %3v %v -> %v", i, max, bracketSequence, nextBracketSequcence))
		bracketSequence = nextBracketSequcence
	}

}
