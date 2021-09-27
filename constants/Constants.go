package constants

import (
	"fmt"
	"strings"
)

const constantIterationIndex = 2147483648 + 1
const constantPreviousIterationValue = 2147483648 + 2
const constantArgument = 2147483648 + 3

func nextCombinationDigit(constants *[]float64,
	combination []float64,
	depth int,
	maxIterationIndexes int,
	maxPreviousIterationValue int,
	maxArguments int,
	previousIterationRequired bool,
	ready func(constantsCombination []float64)) {

	if len(combination) < depth {
		for i := 0; i < len(*constants); i++ {

			v := (*constants)[i]

			currentIterationIndexes := maxIterationIndexes
			currentPreviousIterationValue := maxPreviousIterationValue
			currentArguments := maxArguments

			if v == constantIterationIndex {
				currentIterationIndexes--
			}

			if v == constantPreviousIterationValue {
				currentPreviousIterationValue--
			}

			if v == constantArgument {
				currentArguments--
			}

			if currentIterationIndexes >= 0 && currentPreviousIterationValue >= 0 && currentArguments >= 0 {
				nextCombination := append(combination, v)
				nextCombinationDigit(
					constants,
					nextCombination,
					depth,
					currentIterationIndexes,
					currentPreviousIterationValue,
					currentArguments,
					previousIterationRequired,
					ready)
			}
		}
	} else {
		if !previousIterationRequired || contains(combination, constantPreviousIterationValue) {
			ready(combination)
		}
	}
}

func contains(s []float64, e float64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Recombination recombine all constants and call ready function for each combination
func Recombination(
	constants *[]float64,
	digits int,
	maxIterationIndexes int,
	maxPreviousIterationValue int,
	maxArguments int,
	previousIterationRequired bool,
	ready func(constantsCombination []float64)) {

	combination := []float64{}
	nextCombinationDigit(
		constants,
		combination,
		digits,
		maxIterationIndexes,
		maxPreviousIterationValue,
		maxArguments,
		previousIterationRequired,
		ready)

}

// ToString converts constant to string representation
func ToString(constant float64) string {
	switch constant {
	case constantIterationIndex:
		return "    n"
	case constantPreviousIterationValue:
		return "prevX"
	case constantArgument:
		return "    x"
	default:
		return fmt.Sprintf("%5v", constant)
	}
}

// CombinationToString converts combination of constants to string
func CombinationToString(combination []float64, separator string) string {
	out := []string{}
	for _, v := range combination {
		out = append(out, ToString(v))
	}

	return strings.Join(out, separator)
}
