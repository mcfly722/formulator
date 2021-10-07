package vm

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/mcfly722/formulator/catalan"
)

// go test -v constants.go constants_test.go

func Test_Compilation(t *testing.T) {
	sequence := "(((()()((()))(()())))())"

	bracketSteps, _, _ := catalan.StringToBracketsSteps(sequence)

	fmt.Println(catalan.BracketsStepsToString(bracketSteps))

	a, _ := json.MarshalIndent(bracketSteps, "", "   ")
	fmt.Println(string(a))

	/*
		program, err := Compile(sequence)
		if err != nil {
			t.Errorf("error for %v :%v", sequence, err)
		}
		t.Log(fmt.Sprintf("sequence:%v\n%v", sequence, program.ToString()))
	*/
}
