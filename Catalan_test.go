package main

import (
	"fmt"
	"testing"
)

func nextBracket(tail []BracketStep, sizeX int, sizeY int, output *[]string) {

	if sizeX == 0 && sizeY == 0 {
		*output = append(*output, bracketsStepsToString(tail))
		return
	}

	for i := 1; i <= sizeX; i++ {
		for j := 1; j <= sizeY-(sizeX-i); j++ {
			nextTail := append(tail, BracketStep{x: i, y: j})
			nextBracket(nextTail, sizeX-i, sizeY-j, output)
		}
	}
}

func Test_Sequence(t *testing.T) {
	solutions1 := []string{}
	solutions2 := []string{}

	nextBracket([]BracketStep{}, 8, 8, &solutions1)

	currentSequence := solutions1[0]
	for i := 0; i < len(solutions1); i++ {
		next, err := GetNextBracketsTree(currentSequence)
		if err != nil {
			t.Errorf("getNextBracketsTree('%v') returned error:%v", currentSequence, err)
		}
		solutions2 = append(solutions2, currentSequence)
		currentSequence = next
	}

	for i, solution := range solutions1 {
		t.Log(fmt.Sprintf("%v  %v %v", i+1, solution, solutions2[i]))
		if solution != solutions2[i] {
			t.Errorf("sequences does not match")
		}
	}
}

func testSequenceError(t *testing.T, testSequence string) {
	_, err := GetNextBracketsTree(testSequence)

	if err == nil {
		t.Errorf("error for '%v' sequence does not catched", testSequence)
	} else {
		t.Log(fmt.Sprintf("error for '%v' successfully catched.\nerror description:%v", testSequence, err))
	}
}

func testSequenceSuccess(t *testing.T, testSequence string, expectingAnswer string) {
	answer, err := GetNextBracketsTree(testSequence)

	if err != nil {
		t.Errorf("error for %v :%v", testSequence, err)
	}

	if answer != expectingAnswer {
		t.Errorf("expecting answer %v for %v are not equal to returned answer %v", expectingAnswer, testSequence, answer)
	}

}

func Test_Empty(t *testing.T) {
	testSequenceError(t, "")
}

func Test_IncorrectSymbolFromStart(t *testing.T) {
	testSequenceError(t, "a(()(()))")
}

func Test_IncorrectSymbolAtTheEnd(t *testing.T) {
	testSequenceError(t, "(()(()))b")
}

func Test_IncorrectSymbolInTheMiddle(t *testing.T) {
	testSequenceError(t, "(()((c)))")
}

func Test_LostOpeningBracket(t *testing.T) {
	testSequenceError(t, "())")
}

func Test_LostClosingBracket(t *testing.T) {
	testSequenceError(t, "(()")
}

func Test_WrongBracketsSequence(t *testing.T) {
	testSequenceError(t, "())(")
}

func Test_FirstBracket(t *testing.T) {
	testSequenceSuccess(t, "()", "()()")
}

func Test_FirstNBrackets(t *testing.T) {
	currentSequence := "()"
	for i := 0; i < 120; i++ {
		next, err := GetNextBracketsTree(currentSequence)
		if err != nil {
			t.Errorf("getNextBracketsTree('%v') returned error:%v", currentSequence, err)
		}
		t.Log(fmt.Sprintf("%v %v -> %v", i, currentSequence, next))
		currentSequence = next
	}
}
