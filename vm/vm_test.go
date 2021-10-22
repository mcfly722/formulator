package vm

import (
	"fmt"
	"os"
	"testing"

	"github.com/mcfly722/formulator/constants"
	"github.com/mcfly722/formulator/functions"
	"github.com/mcfly722/formulator/operators"
)

const testBracketSequence = "(())((()()))"

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

func Test_CompilationRecombination(t *testing.T) {
	i := 1

	readyProgram := func(program *Program) {
		fmt.Println(fmt.Sprintf("%3v) %v          %v     %v     %v", i, constants.CombinationOfPointersToString(&program.Constants, " "), operators.CombinationToString(&program.Operators, " "), functions.CombinationToString(&program.Functions, " "), Decompile(program)))
		i++
		if i > 3000 {
			os.Exit(0)
		}
	}

	err := RecombineSequence(testBracketSequence, &constants.AvailableConstants, functions.Functions, operators.Operators, readyProgram)
	if err != nil {
		t.Errorf("%v", err)
	}
}
