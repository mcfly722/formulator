package combinator

import (
	"fmt"
	"os"
	"testing"

	"github.com/mcfly722/formulator/constants"
	"github.com/mcfly722/formulator/functions"
	"github.com/mcfly722/formulator/operators"
	"github.com/mcfly722/formulator/vm"
)

func testRecombination(t *testing.T, testBracketSequence string) {
	t.Log(fmt.Sprintf("sequence: %v", testBracketSequence))

	i := 1
	readyProgram := func(program *vm.Program) {
		t.Log(fmt.Sprintf("%3v) %v          %v     %v     %v", i, constants.CombinationOfPointersToString(&program.Constants, " "), operators.CombinationToString(&program.Operators, " "), functions.CombinationToString(&program.Functions, " "), vm.Decompile(program)))
		i++
		if i > 30 {
			os.Exit(0)
		}
	}

	err := RecombineSequence(testBracketSequence, &constants.AvailableConstants, functions.Functions, operators.Operators, readyProgram)
	if err != nil {
		t.Errorf("%v", err)
	}

}

func Test_Recombination1(t *testing.T) {
	testRecombination(t, "()")
}

func Test_Recombination2(t *testing.T) {
	testRecombination(t, "(())")
}

func Test_Recombination3(t *testing.T) {
	testRecombination(t, "((()))")
}

/*

func Test_Recombination4(t *testing.T) {
	testRecombination(t, "()()")
}

func Test_Recombination5(t *testing.T) {
	testRecombination(t, "(()())")
}
*/

/*
func Test_Recombination6(t *testing.T) {
	testRecombination(t, "((()()))")
}
*/
