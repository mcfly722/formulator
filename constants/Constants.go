package constants

import (
	"fmt"
	"strings"
)

// IterationIndex constant
const IterationIndex = 2147483648 + 1

// PreviousIterationValue constant
const PreviousIterationValue = 2147483648 + 2

// Argument constant
const Argument = 2147483648 + 3

// AvailableConstants all available constants
var AvailableConstants = []float64{3, IterationIndex, PreviousIterationValue, Argument}

func nextCombinationDigit(
	availableConstants *[]float64,
	combination []float64,
	depth int,
	maxIterationIndexes int,
	maxPreviousIterationValue int,
	maxArguments int,
	previousIterationRequired bool,
	constantsCombination *[]*float64,
	ready func(constantsCombination *[]*float64)) {

	if len(combination) < depth {
		for i := 0; i < len(*availableConstants); i++ {

			v := (*availableConstants)[i]

			currentIterationIndexes := maxIterationIndexes
			currentPreviousIterationValue := maxPreviousIterationValue
			currentArguments := maxArguments

			if v == IterationIndex {
				currentIterationIndexes--
			}

			if v == PreviousIterationValue {
				currentPreviousIterationValue--
			}

			if v == Argument {
				currentArguments--
			}

			if currentIterationIndexes >= 0 && currentPreviousIterationValue >= 0 && currentArguments >= 0 {
				nextCombination := append(combination, v)
				nextCombinationDigit(
					availableConstants,
					nextCombination,
					depth,
					currentIterationIndexes,
					currentPreviousIterationValue,
					currentArguments,
					previousIterationRequired,
					constantsCombination,
					ready)
			}
		}
	} else {
		if !previousIterationRequired || contains(combination, PreviousIterationValue) {

			for i, c := range combination {
				*((*constantsCombination)[i]) = c
			}
			ready(constantsCombination)
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
	availableConstants *[]float64,
	constantsCombination *[]*float64,
	maxIterationIndexes int,
	maxPreviousIterationValue int,
	maxArguments int,
	previousIterationRequired bool,
	ready func(constantsCombination *[]*float64)) {

	combination := []float64{}

	nextCombinationDigit(
		availableConstants,
		combination,
		len(*constantsCombination),
		maxIterationIndexes,
		maxPreviousIterationValue,
		maxArguments,
		previousIterationRequired,
		constantsCombination,
		ready)

}

// ToString converts constant to string representation
func ToString(constant float64) string {
	switch constant {
	case IterationIndex:
		return "    n"
	case PreviousIterationValue:
		return "prevX"
	case Argument:
		return "    x"
	default:
		return fmt.Sprintf("%v", constant)
	}
}

// CombinationToString converts combination of constants to string
func CombinationToString(combination *[]float64, separator string) string {
	out := []string{}
	for _, constant := range *combination {
		out = append(out, ToString(constant))
	}

	return strings.Join(out, separator)
}
