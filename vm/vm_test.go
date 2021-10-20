package vm

import (
	"fmt"
	"testing"

	"github.com/mcfly722/formulator/constants"
	"github.com/mcfly722/formulator/functions"
	"github.com/mcfly722/formulator/operators"
)

// go test -v constants.go constants_test.go

func Test_Compilation(t *testing.T) {
	sequence := "(())((()()))"

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

func recombineSequence(sequence string, availableConstants *[]float64, availableFunctions []*functions.Function, availableOperators []*operators.Operator) error {

	program, err := Compile(sequence)
	if err != nil {
		return err
	}

	i := 1

	ready := func(constantsCombination *[]float64) {
		fmt.Println(fmt.Sprintf("%3v) %v", i, constants.CombinationToString(program.Constants, " ")))
		i++
	}

	constants.Recombination(availableConstants, program.Constants, 1, 2, 3, true, ready)

	return nil
}

func Test_CompilationRecombination(t *testing.T) {
	sequence := "()((()))"

	err := recombineSequence(sequence, &constants.AvailableConstants, functions.Functions, operators.Operators)
	if err != nil {
		t.Errorf("%v", err)
	}
}
