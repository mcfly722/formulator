package constants

const constantIterationIndex = 0
const constantPreviousIterationValue = 1
const constantArgument = 2

func nextCombinationDigit(constants *[]float64,
	combination []float64,
	depth int,
	maxIterationIndexes int,
	maxPreviousIterationValue int,
	maxArguments int,
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
					ready)
			}
		}
	} else {
		ready(combination)
	}
}

// Recombination recombine all constants and call ready function for each combination
func Recombination(
	constants *[]float64,
	digits int,
	maxIterationIndexes int,
	maxPreviousIterationValue int,
	maxArguments int,
	ready func(constantsCombination []float64)) {

	combination := []float64{}
	nextCombinationDigit(
		constants,
		combination,
		digits,
		maxIterationIndexes,
		maxPreviousIterationValue,
		maxArguments,
		ready)

}
