package catalan

import (
	"encoding/json"
	"fmt"
	"testing"
)

// go test -v

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
		next, err := GetNextBracketsSequence(currentSequence)
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

func testGetNextBracketsSequenceError(t *testing.T, testSequence string) {
	_, err := GetNextBracketsSequence(testSequence)

	if err == nil {
		t.Errorf("error for '%v' sequence does not catched", testSequence)
	} else {
		t.Log(fmt.Sprintf("error for '%v' successfully catched.\nerror description:%v", testSequence, err))
	}
}

func testGetNextBracketsSequenceSuccess(t *testing.T, testSequence string, expectingAnswer string) {
	answer, err := GetNextBracketsSequence(testSequence)

	if err != nil {
		t.Errorf("error for %v :%v", testSequence, err)
	}

	if answer != expectingAnswer {
		t.Errorf("expecting answer %v for %v are not equal to returned answer %v", expectingAnswer, testSequence, answer)
	}

}

func Test_Empty(t *testing.T) {
	testGetNextBracketsSequenceError(t, "")
}

func Test_GetNextBracketsSequence_IncorrectSymbolFromStart(t *testing.T) {
	testGetNextBracketsSequenceError(t, "a(()(()))")
}

func Test_GetNextBracketsSequence_IncorrectSymbolAtTheEnd(t *testing.T) {
	testGetNextBracketsSequenceError(t, "(()(()))b")
}

func Test_GetNextBracketsSequence_IncorrectSymbolInTheMiddle(t *testing.T) {
	testGetNextBracketsSequenceError(t, "(()((c)))")
}

func Test_GetNextBracketsSequence_LostOpeningBracket(t *testing.T) {
	testGetNextBracketsSequenceError(t, "())")
}

func Test_GetNextBracketsSequence_LostClosingBracket(t *testing.T) {
	testGetNextBracketsSequenceError(t, "(()")
}

func Test_GetNextBracketsSequence_WrongBracketsSequence(t *testing.T) {
	testGetNextBracketsSequenceError(t, "())(")
}

func Test_GetNextBracketsSequence_FirstBracket(t *testing.T) {
	testGetNextBracketsSequenceSuccess(t, "()", "()()")
}

func Test_GetNextBracketsSequence_FirstNBrackets(t *testing.T) {
	currentSequence := "()"
	for i := 0; i < 120; i++ {
		next, err := GetNextBracketsSequence(currentSequence)
		if err != nil {
			t.Errorf("getNextBracketsTree('%v') returned error:%v", currentSequence, err)
		}
		t.Log(fmt.Sprintf("%v %v -> %v", i, currentSequence, next))
		currentSequence = next
	}
}

func Test_BracketsToExpressionTree(t *testing.T) {
	bracketSequence := "()((())())"
	expression, err := BracketsToExpressionTree(bracketSequence)
	if err != nil {
		t.Errorf("Cant build expression tree for %v. Error: %v", bracketSequence, err)
	}

	bytes, err := json.Marshal(expression)
	if err != nil {
		t.Errorf("Can't serialize %v. Error:%v", bracketSequence, err)
	}
	t.Log(string(bytes))

}

func testBracketsToExpressionTreeError(t *testing.T, testSequence string) {
	_, err := BracketsToExpressionTree(testSequence)

	if err == nil {
		t.Errorf("error for '%v' sequence does not catched", testSequence)
	} else {
		t.Log(fmt.Sprintf("error for '%v' successfully catched.\nerror description:%v", testSequence, err))
	}
}

func testBracketsToExpressionTreeSuccess(t *testing.T, testSequence string) {
	expression, err := BracketsToExpressionTree(testSequence)

	if err != nil {
		t.Errorf("error for %v :%v", testSequence, err)
	}

	bytes, err := json.Marshal(expression)
	if err != nil {
		t.Errorf("Can't serialize %v. Error:%v", testSequence, err)
	}
	t.Log(string(bytes))
}

func Test_BracketsToExpressionTree_IncorrectSymbolFromStart(t *testing.T) {
	testBracketsToExpressionTreeError(t, "a(()(()))")
}

func Test_BracketsToExpressionTree_IncorrectSymbolAtTheEnd(t *testing.T) {
	testBracketsToExpressionTreeError(t, "(()(()))b")
}

func Test_BracketsToExpressionTree_IncorrectSymbolInTheMiddle(t *testing.T) {
	testBracketsToExpressionTreeError(t, "(()((c)))")
}

func Test_BracketsToExpressionTree_LostOpeningBracket(t *testing.T) {
	testBracketsToExpressionTreeError(t, "())")
}

func Test_BracketsToExpressionTree_LostClosingBracket(t *testing.T) {
	testBracketsToExpressionTreeError(t, "(()")
}

func Test_BracketsToExpressionTree_WrongBracketsSequence(t *testing.T) {
	testBracketsToExpressionTreeError(t, "())(")
}

func Test_BracketsToExpressionTree_FirstBracket(t *testing.T) {
	testBracketsToExpressionTreeSuccess(t, "()")
}

func Test_NewCombination_success(t *testing.T) {
	n := 2
	k := 5
	sequence, err := NewCombination(n, k)
	if err != nil {
		t.Errorf("Error:%v", err)
		return
	}
	if len(sequence) != k {
		t.Errorf(fmt.Sprintf("wrong sequence len(%v) != %v", sequence, k))
		return
	}

	t.Log(fmt.Sprintf("new combination %v from %v is %v len=%v is correct", n, k, sequence, len(sequence)))
}

func Test_CombinationNKNext(t *testing.T) {
	currentSequence, err := NewCombination(3, 8)

	if err != nil {
		t.Errorf("Error:%v", err)
		return
	}

	finished := false
	for i := 1; finished == false; i++ {
		t.Log(fmt.Sprintf("%2v %v = %v", i, currentSequence, Combination2Int(currentSequence)))

		combination, finishedTmp, err := CombinationNKNext(currentSequence)

		if err != nil {
			t.Errorf("Error:%v", err)
			return
		}

		finished = finishedTmp

		currentSequence = combination
	}

}

// go test -bench .

func BenchmarkGetNextBracketsSequence(b *testing.B) {
	currentSequence := "()"
	for n := 0; n < b.N; n++ {
		next, _ := GetNextBracketsSequence(currentSequence)
		currentSequence = next
	}
}

func BenchmarkBracketsToExpressionTree(b *testing.B) {
	currentSequence := "()((())())((()((())())))"
	for n := 0; n < b.N; n++ {
		_, err := BracketsToExpressionTree(currentSequence)
		if err != nil {
			b.Errorf("error:%v", err)
		}
	}
}
