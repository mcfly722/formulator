package vm

import (
	"encoding/json"
	"fmt"

	"github.com/mcfly722/formulator/functions"
	"github.com/mcfly722/formulator/operators"
	"github.com/mcfly722/formulator/zeroOneTwoTree"
)

const function = 1
const operator = 2

// Instruction represents virtual machine instruction
type Instruction struct {
	operand1 *float64
	operand2 *float64
	operator operators.Operator
	function functions.Function
	kind     int
	result   float64
}

// Program represents program for execution
type Program struct {
	Constants       []float64
	Code            []Instruction
	Functions       []*functions.Function
	Operators       []*operators.Operator
	ASTTree         *zeroOneTwoTree.Node
	BracketSequence string
}

func hiveCompilation(node *zeroOneTwoTree.Node, program *Program) *float64 {

	// constant
	if len(node.Childs) == 0 {
		var constant float64
		program.Constants = append(program.Constants, constant)
		//fmt.Println("constant")
		return &constant
	}

	// function
	if len(node.Childs) == 1 {
		instruction := Instruction{
			operand1: hiveCompilation(node.Childs[0], program),
			kind:     function,
		}
		program.Functions = append(program.Functions, &instruction.function)
		program.Code = append(program.Code, instruction)
		return &instruction.result
	}

	// operator
	instruction := Instruction{
		operand1: hiveCompilation(node.Childs[0], program),
		operand2: hiveCompilation(node.Childs[1], program),
		kind:     operator,
	}

	program.Operators = append(program.Operators, &instruction.operator)
	program.Code = append(program.Code, instruction)

	return &instruction.result
}

// Compile brackets string to Program
func Compile(bracketsSequence string) (*Program, error) {

	astTree, err := zeroOneTwoTree.BracketsToTree(bracketsSequence)
	if err != nil {
		return nil, err
	}

	program := &Program{
		Constants:       []float64{},
		Code:            []Instruction{},
		Functions:       []*functions.Function{},
		Operators:       []*operators.Operator{},
		ASTTree:         astTree,
		BracketSequence: bracketsSequence,
	}

	hiveCompilation(astTree, program)

	return program, nil
}

// ToString represents program string debug information
func (program *Program) ToString() string {
	astJSON, _ := json.MarshalIndent(program.ASTTree, "", "  ")
	return fmt.Sprintf("program\nsequence:%v\nconstants:%v\nfunctions:%v\noperators:%v\ninstructions:%v\nASTTree:\n%v", program.BracketSequence, len(program.Constants), len(program.Functions), len(program.Operators), len(program.Code), string(astJSON))
}
