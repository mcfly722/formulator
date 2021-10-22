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

/*
func Test_CompilationError1(t *testing.T) {
	sequence := "(()()()!)"
	_, err := Compile(sequence)
	if err != nil {
		t.Log(fmt.Sprintf("Correct handling error for %v", sequence))
	} else {
		t.Errorf("Incorrect compilation for %v", sequence)
	}
}
*/

func Test_Decompilation(t *testing.T) {
	sequence := testBracketSequence
	program, err := Compile(sequence)
	if err != nil {
		t.Errorf("Compilation ('%v') returned error: %v", sequence, err)
	}

	t.Log(fmt.Sprintf("decompiled to: %v", Decompile(program)))
}

func recombineSequence(sequence string, availableConstants *[]float64, availableFunctions []*functions.Function, availableOperators []*operators.Operator) error {

	program, err := Compile(sequence)
	if err != nil {
		return err
	}

	i := 1

	readyConstants := func(constantsCombination *[]*float64) {

		if len(program.Operators) > 0 {

			readyOperators := func(operatorsCombination []*operators.Operator) {

				if len(program.Functions) > 0 {
					readyFunctions := func(functionsCombination []*functions.Function) {
						fmt.Println(fmt.Sprintf("%3v) %v          %v     %v     %v", i, constants.CombinationOfPointersToString(&program.Constants, " "), operators.CombinationToString(&program.Operators, " "), functions.CombinationToString(&program.Functions, " "), Decompile(program)))
						i++
						if i > 3000 {
							os.Exit(0)
						}
					}

					functions.Recombination(availableFunctions, program.Functions, readyFunctions)
					fmt.Println("")
				}
			}

			operators.Recombination(availableOperators, program.Operators, readyOperators)
			fmt.Println("")
		}

	}

	constants.Recombination(availableConstants, &program.Constants, 1, 2, 3, true, readyConstants)

	return nil
}

func Test_CompilationRecombination(t *testing.T) {
	err := recombineSequence(testBracketSequence, &constants.AvailableConstants, functions.Functions, operators.Operators)
	if err != nil {
		t.Errorf("%v", err)
	}
}
