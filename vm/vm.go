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
	operand1 *float64
	operand2 *float64
	operator operators.Operator
	function functions.Function
	kind     int
	result   float64
}

// Program represents program for execution
type Program struct {
	Constants       []*float64
	Code            []*Instruction
	Functions       []*functions.Function
	Operators       []*operators.Operator
	ASTTree         *zeroOneTwoTree.Node
	BracketSequence string
}

func childInstructionByResultPointer(instructions []*Instruction, resultPointer *float64) (int, bool) {
	for i, instruction := range instructions {
		if resultPointer == &instruction.result {
			return i, true
		}
	}
	return 0, false
}

func decompileHive(instructions []*Instruction, i int) string {

	var instruction = *instructions[i]
	//	fmt.Println(fmt.Sprintf("instruction #%v kind:%v i_addr=%p 1op_addr:%p 2op_addr:%p result_addr:%p", i, instruction.kind, instructions[i], instruction.operand1, instruction.operand2, &instruction.result))

	operand1 := "unknown operand 1"
	operand2 := "unknown operand 2"

	i1, found1 := childInstructionByResultPointer(instructions, instruction.operand1)
	//	fmt.Println(fmt.Sprintf("%v", found1))
	if found1 {
		operand1 = decompileHive(instructions, i1)
	} else {
		if instruction.operand1 != nil {
			operand1 = constants.ToString(*instruction.operand1)
		}
	}

	i2, found2 := childInstructionByResultPointer(instructions, instruction.operand2)

	//	fmt.Println(fmt.Sprintf("%v", found2))
	if found2 {
		operand2 = decompileHive(instructions, i2)
	} else {
		if instruction.operand2 != nil {
			operand2 = constants.ToString(*instruction.operand2)
		}
	}

	if instruction.kind == function {
		return fmt.Sprintf("%v[%v]", instruction.function.Name, operand1)
	}

	if instruction.kind == operator {

		//		fmt.Println(fmt.Sprintf("op1  :%p", instruction.operand1))
		//		fmt.Println(fmt.Sprintf("op2  :%p", instruction.operand2))

		return fmt.Sprintf("(%v%v%v)", operand1, instruction.operator.Name, operand2)
	}

	return "unknown instruction"
}

// Decompile program to string representation
func Decompile(program *Program) string {
	return decompileHive(program.Code, len(program.Code)-1)
}

func hiveCompilation(node *zeroOneTwoTree.Node, program *Program) *float64 {

	// constant
	if len(node.Childs) == 0 {
		var constant float64
		program.Constants = append(program.Constants, &constant)
		//fmt.Println(fmt.Sprintf("      constant:%p", &constant))
		return &constant
	}

	// function
	if len(node.Childs) == 1 {
		instruction := Instruction{
			operand1: hiveCompilation(node.Childs[0], program),
			function: functions.Function{Name: "f"},
			kind:     function,
		}
		program.Functions = append(program.Functions, &(instruction.function))
		program.Code = append(program.Code, &instruction)

		//fmt.Println(fmt.Sprintf("f instruction:%p", &(instruction)))
		//fmt.Println(fmt.Sprintf("      operand:%p", instruction.operand1))

		return &instruction.result
	}

	// operator
	if len(node.Childs) == 2 {
		instruction := Instruction{
			operand1: hiveCompilation(node.Childs[0], program),
			operand2: hiveCompilation(node.Childs[1], program),
			operator: operators.Operator{Name: " ? "},
			kind:     operator,
		}

		program.Operators = append(program.Operators, &(instruction.operator))
		program.Code = append(program.Code, &instruction)

		//fmt.Println(fmt.Sprintf("o instruction:%p", &(instruction)))
		//fmt.Println(fmt.Sprintf("      operand:%p", instruction.operand1))
		//fmt.Println(fmt.Sprintf("      operand:%p", instruction.operand2))

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
		Code:            []*Instruction{},
		Functions:       []*functions.Function{},
		Operators:       []*operators.Operator{},
		ASTTree:         astTree,
		BracketSequence: bracketsSequence,
	}

	hiveCompilation(astTree, &program)

	/*
		fmt.Print("constants:")
		for _, p := range program.Constants {
			fmt.Print(fmt.Sprintf("%p ", p))
		}
		fmt.Println("")
	*/

	return &program, nil
}

// ToString represents program string debug information
func (program *Program) ToString() string {
	astJSON, _ := json.MarshalIndent(program.ASTTree, "", "  ")
	return fmt.Sprintf("program\nsequence:%v\nconstants:%v\nfunctions:%v\noperators:%v\ninstructions:%v\nASTTree:\n%v", program.BracketSequence, len(program.Constants), len(program.Functions), len(program.Operators), len(program.Code), string(astJSON))
}
