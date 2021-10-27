package combinator

import (
	"errors"
	"fmt"
	"math"

	"github.com/mcfly722/formulator/constants"
	"github.com/mcfly722/formulator/functions"
	"github.com/mcfly722/formulator/operators"
	"github.com/mcfly722/formulator/vm"
)

// Point struct
type Point struct {
	X float64
	Y float64
}

// NewPoint constructor
func NewPoint(_x float64, _y float64) *Point {
	return &Point{
		X: _x,
		Y: _y,
	}
}

const testBracketSequence = "((()()))()"

// RecombineSequence function
func RecombineSequence(
	sequence string,
	availableConstants *[]float64,
	availableFunctions []*functions.Function,
	availableOperators []*operators.Operator,
	maxIterationIndexes int,
	maxPreviousIterationValue int,
	maxArguments int,
	previousIterationRequired bool,
	readyProgram func(program *vm.Program) bool) error {

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
				continueCalculations := readyProgram(program)
				if !continueCalculations {
					return
				}
			}

		}
	}

	if len(program.Constants) > 0 {
		constants.Recombination(availableConstants, &program.Constants, maxIterationIndexes, maxPreviousIterationValue, maxArguments, previousIterationRequired, readyConstants)
	} else {
		return errors.New("there are no constants to iterate")
	}

	return nil
}

// CalculateDeviation calculate deviation for all points and compare it with threshold
func CalculateDeviation(program *vm.Program, points *[]Point, deviationThreshold float64, numberOfIterations int, debug bool) (float64, error) {
	var deviation float64

	for pi, point := range *points {

		var previousValue0 float64
		var previousValue1 float64 = 1
		var value float64

		for iteration := numberOfIterations - 1; iteration >= 0; iteration-- {
			// replace constants with initial values for current point

			constantsBackup := []float64{}
			for _, constant := range program.Constants {
				constantsBackup = append(constantsBackup, *constant)

				switch *constant {
				case constants.N:
					*constant = float64(iteration)
				case constants.X:
					*constant = point.X
				case constants.PreviousValue0:
					*constant = previousValue0
				case constants.PreviousValue1:
					*constant = previousValue1
				}
			}

			if debug {
				fmt.Println(fmt.Sprintf("constants=%v iteration=%v", constantsBackup, iteration))
			}

			// calculate iteration
			value = program.Calculate()

			// restore from backup
			for j, constant := range program.Constants {
				*constant = constantsBackup[j]
			}

			if debug {
				fmt.Println(fmt.Sprintf("pointN=%v (%v,%v) iteration=%v value=%v", pi, point.X, point.Y, iteration, value))
			}

			if math.IsNaN(value) || math.IsInf(value, 1) || math.IsInf(value, -1) {
				return value, fmt.Errorf("calculation out of rational range")
			}

			previousValue0 = value
			previousValue1 = value

		}

		deviation += math.Abs(value - point.Y)

		if deviation > deviationThreshold {
			return deviation, fmt.Errorf("deviation threshold reached")
		}
	}

	return deviation, nil
}
