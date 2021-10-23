package combinator

import (
	"errors"
	"fmt"

	"github.com/mcfly722/formulator/constants"
	"github.com/mcfly722/formulator/functions"
	"github.com/mcfly722/formulator/operators"
	"github.com/mcfly722/formulator/vm"
)

const testBracketSequence = "((()()))()"

// RecombineSequence function
func RecombineSequence(sequence string, availableConstants *[]float64, availableFunctions []*functions.Function, availableOperators []*operators.Operator, readyProgram func(program *vm.Program)) error {

	program, err := vm.Compile(sequence)
	if err != nil {
		return err
	}

	fmt.Println(program.ToString())

	readyConstants := func(constantsCombination *[]*float64) {

		readyFunctions := func(functionsCombination []*functions.Function) {
			readyProgram(program)
		}

		readyOperators := func(operatorsCombination []*operators.Operator) {
			if len(program.Functions) > 0 {
				functions.Recombination(availableFunctions, program.Functions, readyFunctions)
			} else {
				readyProgram(program)
			}
		}

		if len(program.Operators) > 0 {
			operators.Recombination(availableOperators, program.Operators, readyOperators)
		} else {

			if len(program.Functions) > 0 {
				functions.Recombination(availableFunctions, program.Functions, readyFunctions)
			} else {
				readyProgram(program)
			}

		}
	}

	if len(program.Constants) > 0 {
		constants.Recombination(availableConstants, &program.Constants, 1, 2, 3, false, readyConstants)
	} else {
		return errors.New("there are no constants to iterate")
	}

	return nil
}
