package vm

import (
	"encoding/json"
	"fmt"

	"github.com/mcfly722/formulator/constants"
	"github.com/mcfly722/formulator/functions"
	"github.com/mcfly722/formulator/operators"
	"github.com/mcfly722/formulator/zeroOneTwoTree"
)

const function = 1
const operator = 2

// Instruction represents virtual machine instruction
type Instruction struct {
	operand1  *float64
	operand2  *float64
	operatorN int
	functionN int
	kind      int
	result    float64
}

// Program represents program for execution
type Program struct {
	Constants       []*float64
	Instructions    []*Instruction
	Functions       []*functions.Function
	Operators       []*operators.Operator
	ASTTree         *zeroOneTwoTree.Node
	BracketSequence string
}

// Calculate final program result
func (program *Program) Calculate() float64 {
	var result float64 = -1
	for _, instruction := range program.Instructions {
		if instruction.kind == function {
			result = program.Functions[instruction.functionN].Function(*instruction.operand1)
		}

		if instruction.kind == operator {
			result = program.Operators[instruction.operatorN].Function(*instruction.operand1, *instruction.operand2)
		}

		instruction.result = result
	}
	return result
}

func childInstructionByResultPointer(instructions []*Instruction, resultPointer *float64) (int, bool) {
	for i, instruction := range instructions {
		if resultPointer == &instruction.result {
			return i, true
		}
	}
	return 0, false
}

func decompileHive(program *Program, i int) string {
	var instruction = *program.Instructions[i]
	operand1 := "unknown operand 1"
	operand2 := "unknown operand 2"

	i1, found1 := childInstructionByResultPointer(program.Instructions, instruction.operand1)
	if found1 {
		operand1 = decompileHive(program, i1)
	} else {
		if instruction.operand1 != nil {
			operand1 = constants.ToString(*instruction.operand1)
		}
	}

	i2, found2 := childInstructionByResultPointer(program.Instructions, instruction.operand2)

	if found2 {
		operand2 = decompileHive(program, i2)
	} else {
		if instruction.operand2 != nil {
			operand2 = constants.ToString(*instruction.operand2)
		}
	}

	if instruction.kind == function {
		return fmt.Sprintf("%v[%v]", program.Functions[instruction.functionN].Name, operand1)
	}

	if instruction.kind == operator {
		return fmt.Sprintf("(%v%v%v)", operand1, program.Operators[instruction.operatorN].Name, operand2)
	}

	return "unknown instruction"
}

// Decompile program to string representation
func Decompile(program *Program) string {
	return decompileHive(program, len(program.Instructions)-1)
}

func hiveCompilation(node *zeroOneTwoTree.Node, program *Program) *float64 {

	// constant
	if len(node.Childs) == 0 {
		var constant float64
		program.Constants = append(program.Constants, &constant)
		return &constant
	}

	// function
	if len(node.Childs) == 1 {
		instruction := Instruction{
			operand1:  hiveCompilation(node.Childs[0], program),
			functionN: len(program.Functions),
			kind:      function,
		}
		program.Functions = append(program.Functions, &(functions.Function{Name: "f"}))
		program.Instructions = append(program.Instructions, &instruction)

		return &instruction.result
	}

	// operator
	if len(node.Childs) == 2 {
		instruction := Instruction{
			operand1:  hiveCompilation(node.Childs[0], program),
			operand2:  hiveCompilation(node.Childs[1], program),
			operatorN: len(program.Operators),
			kind:      operator,
		}

		program.Operators = append(program.Operators, &(operators.Operator{Name: " ? "}))
		program.Instructions = append(program.Instructions, &instruction)

		return &instruction.result
	}

	return nil
}

// Compile brackets string to Program
func Compile(bracketsSequence string) (*Program, error) {

	astTree, err := zeroOneTwoTree.BracketsToTree(bracketsSequence)
	if err != nil {
		return nil, err
	}

	program := Program{
		Constants:       []*float64{},
		Instructions:    []*Instruction{},
		Functions:       []*functions.Function{},
		Operators:       []*operators.Operator{},
		ASTTree:         astTree,
		BracketSequence: bracketsSequence,
	}

	hiveCompilation(astTree, &program)

	return &program, nil
}

// ToString represents program string debug information
func (program *Program) ToString() string {
	astJSON, _ := json.MarshalIndent(program.ASTTree, "", "  ")
	return fmt.Sprintf("program\nsequence:%v\nconstants:%v\nfunctions:%v\noperators:%v\ninstructions:%v\nASTTree:\n%v", program.BracketSequence, len(program.Constants), len(program.Functions), len(program.Operators), len(program.Instructions), string(astJSON))
}
