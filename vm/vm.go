package vm

import (
	"encoding/json"
	"fmt"

	"github.com/mcfly722/formulator/catalan"
)

// Instruction represents virtual machine instruction
type Instruction struct {
	operand1 *float64
	operand2 *float64
	process  func(operand1 float64, operand2 float64) float64
	result   float64
}

// Program represents program for execution
type Program struct {
	Constants *[]float64
	Functions *[]*(func(operand1 float64, operand2 float64) float64)
	Operators *[]*(func(operand1 float64, operand2 float64) float64)
	Code      *[]Instruction
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

func hiveCompilation(expression *catalan.Expression, program *Program) *float64 {
	fmt.Print(".")

	// constant
	if len(expression.Arguments) == 0 {
		var constant float64
		*(program.Constants) = append(*(program.Constants), constant)

		fmt.Println("constant")
		return &constant
	}

	// function
	if len(expression.Arguments) == 1 {
		instruction := Instruction{
			operand1: hiveCompilation(expression.Arguments[0], program),
			operand2: nil,
			process:  func(a float64, b float64) float64 { return a - a + b - b },
			result:   0,
		}
		*(program.Functions) = append(*(program.Functions), &instruction.process)
		*(program.Code) = append(*(program.Code), instruction)

		fmt.Println("function")
		return &instruction.result
	}

	// operator
	instruction := Instruction{
		operand1: hiveCompilation(expression.Arguments[0], program),
		operand2: hiveCompilation(expression.Arguments[1], program),
		process:  func(a float64, b float64) float64 { return a - a + b - b },
		result:   0,
	}
	*(program.Operators) = append(*(program.Operators), &instruction.process)
	*(program.Code) = append(*(program.Code), instruction)

	fmt.Println("operator")
	return &instruction.result
}

// Compile brackets string to Program
func Compile(brackets string) (*Program, error) {

	expression, error := catalan.BracketsToExpressionTree(brackets)
	if error != nil {
		return nil, error
	}

	a, _ := json.MarshalIndent(expression, "", "   ")
	fmt.Println(string(a))

	constants := []float64{}
	functions := []*(func(operand1 float64, operand2 float64) float64){}
	operators := []*(func(operand1 float64, operand2 float64) float64){}
	code := []Instruction{}

	program := &Program{
		Constants: &constants,
		Functions: &functions,
		Operators: &operators,
		Code:      &code,
	}

	hiveCompilation(expression, program)

	return program, nil
}

// ToString represents program string debug information
func (program *Program) ToString() string {
	return fmt.Sprintf("program\n  constants:%v\n  functions:%v\n  operators:%v\n  instructions:%v\n", len(*program.Constants), len(*program.Functions), len(*program.Operators), len(*program.Code))
}
