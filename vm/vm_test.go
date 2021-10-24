package vm

import (
	"fmt"
	"testing"
)

const testBracketSequence = "((()()))()"

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

func testDecompilation(t *testing.T, sequence string) {
	program, err := Compile(sequence)
	if err != nil {
		t.Errorf("Compilation ('%v') returned error: %v", sequence, err)
	}

	decompiled, err2 := Decompile(program)
	if err2 != nil {
		t.Errorf("Decompilation error for '%v': %v", sequence, err2)
	}

	t.Log(fmt.Sprintf("decompiled to: %v", decompiled))
}

func testDecompilationError(t *testing.T, sequence string) {
	program, err := Compile(sequence)
	if err == nil {
		t.Errorf("Could not catch error for %v", sequence)
	}

	decompiled, err := Decompile(program)
	if err == nil {
		t.Errorf("Could not catch error for %v, decompiled to:%v", sequence, decompiled)
	}
	t.Log(fmt.Sprintf("Successfully catched decompilation error for %v: %v", sequence, err))

}

func Test_Decompilation1(t *testing.T) {
	testDecompilationError(t, "")
}

func Test_Decompilation2(t *testing.T) {
	testDecompilation(t, "()")
}
func Test_Decompilation3(t *testing.T) {
	testDecompilation(t, "((()()))()")
}
