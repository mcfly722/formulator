package vm

import (
	"encoding/json"
	"fmt"

	"github.com/mcfly722/formulator/functions"
	"github.com/mcfly722/formulator/operators"
	"github.com/mcfly722/formulator/zeroOneTwoTree"
)

// Instruction represents virtual machine instruction
type Instruction struct {
	operand1 *float64
	operand2 *float64
	process  func(operand1 float64, operand2 float64) float64
	debug    func(operand1 float64, operand2 float64) string
	result   float64
}

// Program represents program for execution
type Program struct {
	Constants       *[]float64
	Functions       *[]*(func(operand1 float64, operand2 float64) float64)
	Operators       *[]*(func(operand1 float64, operand2 float64) float64)
	Code            *[]Instruction
	ASTTree         *zeroOneTwoTree.Node
	BracketSequence string
}

// Execute method to execute all program instructions
func (program *Program) Execute() float64 {
	var result float64
	for i := range *program.Code {
		instruction := (*program.Code)[i]
		r := instruction.process(*instruction.operand1, *instruction.operand2)
		result = r
	}
	return result
}

// Debug program
func (program *Program) Debug(functions []*functions.Function, operators []*operators.Operator) string {
	log := ""
	var result float64
	for i := range *program.Code {
		instruction := (*program.Code)[i]
		r := instruction.process(*instruction.operand1, *instruction.operand2)

		log += fmt.Sprintf("%v) %v\n", i, instruction.debug(*instruction.operand1, *instruction.operand2))

		result = r
	}
	log += fmt.Sprintf("result=%v", result)
	return log
}

func hiveCompilation(node *zeroOneTwoTree.Node, program *Program) *float64 {

	// constant
	if len(node.Childs) == 0 {
		var constant float64
		*(program.Constants) = append(*(program.Constants), constant)
		//fmt.Println("constant")
		return &constant
	}

	// function
	if len(node.Childs) == 1 {
		instruction := Instruction{
			operand1: hiveCompilation(node.Childs[0], program),
			operand2: nil,
			process:  func(a float64, b float64) float64 { return a - a + b - b },
			debug:    func(a float64, b float64) string { return fmt.Sprintf("?(%v)", a) },
			result:   0,
		}
		*(program.Functions) = append(*(program.Functions), &instruction.process)
		*(program.Code) = append(*(program.Code), instruction)

		return &instruction.result
	}

	// operator
	instruction := Instruction{
		operand1: hiveCompilation(node.Childs[0], program),
		operand2: hiveCompilation(node.Childs[1], program),
		process:  func(a float64, b float64) float64 { return a - a + b - b },
		debug:    func(a float64, b float64) string { return fmt.Sprintf("(%v & %v)", a, b) },
		result:   0,
	}
	*(program.Operators) = append(*(program.Operators), &instruction.process)
	*(program.Code) = append(*(program.Code), instruction)

	return &instruction.result
}

// Compile brackets string to Program
func Compile(bracketsSequence string) (*Program, error) {

	astTree, err := zeroOneTwoTree.BracketsToTree(bracketsSequence)
	if err != nil {
		return nil, err
	}

	constants := []float64{}
	functions := []*(func(operand1 float64, operand2 float64) float64){}
	operators := []*(func(operand1 float64, operand2 float64) float64){}
	code := []Instruction{}

	program := &Program{
		Constants:       &constants,
		Functions:       &functions,
		Operators:       &operators,
		Code:            &code,
		ASTTree:         astTree,
		BracketSequence: bracketsSequence,
	}

	hiveCompilation(astTree, program)

	return program, nil
}

// ToString represents program string debug information
func (program *Program) ToString() string {
	astJSON, _ := json.MarshalIndent(program.ASTTree, "", "  ")
	return fmt.Sprintf("program\nsequence:%v\nconstants:%v\nfunctions:%v\noperators:%v\ninstructions:%v\nASTTree:\n%v", program.BracketSequence, len(*program.Constants), len(*program.Functions), len(*program.Operators), len(*program.Code), string(astJSON))
}
