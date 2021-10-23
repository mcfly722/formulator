package vm

import (
	"fmt"
	"testing"
)

const testBracketSequence = "((()()))()"

// go test -v constants.go constants_test.go

func Test_Compilation(t *testing.T) {
	sequence := testBracketSequence
	program, err := Compile(sequence)
	if err != nil {
		t.Errorf("Compilation ('%v') returned error: %v", sequence, err)
	}
	t.Log(fmt.Sprintf("%v", program.ToString()))
}

func Test_CompilationError1(t *testing.T) {
	sequence := "(()()()!)"
	_, err := Compile(sequence)
	if err != nil {
		t.Log(fmt.Sprintf("Correct handling error for %v", sequence))
	} else {
		t.Errorf("Incorrect compilation for %v", sequence)
	}
}

func Test_Decompilation(t *testing.T) {
	sequence := testBracketSequence
	program, err := Compile(sequence)
	if err != nil {
		t.Errorf("Compilation ('%v') returned error: %v", sequence, err)
	}

	t.Log(fmt.Sprintf("decompiled to: %v", Decompile(program)))
}
