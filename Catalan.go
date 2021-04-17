package catalan

import "fmt"

// BracketStep node of binary tree
type BracketStep struct {
	x int
	y int
}

// Expression from bracket
type Expression struct {
	Arguments []*Expression
}

func stringToBracketsSteps(input string) ([]BracketStep, int, error) {
	if input == "" {
		return nil, -1, fmt.Errorf("could not parse empty string, use () for first sequence instead")
	}

	steps := []BracketStep{}
	i, x, y := 0, 0, 0

	counter := 0

	for i < len(input) {
		step := BracketStep{
			x: 0, y: 0,
		}

		for ; i < len(input) && input[i] == '('; i++ {
			step.x++
			x++
			counter++
		}

		for ; i < len(input) && input[i] == ')'; i++ {
			step.y++
			y++
			counter--
			if counter < 0 {
				return nil, -1, fmt.Errorf("%v<- incorrect brackets balance, could not close not opened bracket", input[0:i+1])
			}
		}

		if i < len(input) && input[i] != '(' && input[i] != ')' {
			return nil, -1, fmt.Errorf("%v<- unknown symbol", input[0:i+1])
		}
		steps = append(steps, step)
	}

	if x != y {
		return nil, -1, fmt.Errorf("number of opened brackets (%v) are not equal to closed (%v)", x, y)
	}
	return steps, x, nil
}

func getNextStepsTail(srcTail []BracketStep, dstTail []BracketStep, sizeX int, sizeY int, previousSolutionAlreadyReached bool) ([]BracketStep, bool, bool) {

	if sizeX == 0 && sizeY == 0 {
		if !previousSolutionAlreadyReached {
			return []BracketStep{}, true, false
		}
		return dstTail, true, true
	}

	newDstTail := []BracketStep{}

	startX := 1

	if !previousSolutionAlreadyReached && len(srcTail) > 0 {
		startX = srcTail[0].x
	}

	for i := startX; i <= sizeX; i++ {

		startY := 1
		if !previousSolutionAlreadyReached && len(srcTail) > 0 {
			startY = srcTail[0].y
		}

		for j := startY; j <= sizeY-(sizeX-i); j++ {

			newDstTail = append(dstTail, BracketStep{x: i, y: j})

			newSrcTail := []BracketStep{}
			if len(srcTail) > 0 {
				newSrcTail = srcTail[1:len(srcTail)]
			}

			tail, reached, solutionFound := getNextStepsTail(newSrcTail, newDstTail, sizeX-i, sizeY-j, previousSolutionAlreadyReached)

			previousSolutionAlreadyReached = reached

			if solutionFound {
				return tail, previousSolutionAlreadyReached, solutionFound
			}

		}
	}

	return []BracketStep{}, previousSolutionAlreadyReached, false
}

// GetNextBracketsSequence generates next brackets tree sequence based on previous one from input
func GetNextBracketsSequence(input string) (string, error) {

	tail, size, err := stringToBracketsSteps(input)
	if err != nil {
		return "", err
	}

	nextTail, _, solutionFound := getNextStepsTail(tail, []BracketStep{}, size, size, false)

	if !solutionFound {
		output := ""
		for i := 0; i < size+1; i++ {
			output += "()"
		}
		return output, nil
	}

	return bracketsStepsToString(nextTail), nil
}

func bracketsStepsToString(tail []BracketStep) string {
	output := ""
	for _, step := range tail {
		for i := 0; i < step.x; i++ {
			output += fmt.Sprintf("(")
		}
		for i := 0; i < step.y; i++ {
			output += fmt.Sprintf(")")
		}
	}
	return output
}

// BracketsToExpressionTree generates expression tree based on string of brackets
func BracketsToExpressionTree(input string) (*Expression, error) {
	root := Expression{Arguments: []*Expression{}}

	if input == "" {
		return &root, nil
	}

	counter := 0
	from := 0

	for i := 0; i < len(input); i++ {

		if input[i] == '(' {
			if counter == 0 {
				from = i
			}
			counter++
		}

		if input[i] == ')' {
			counter--

			if counter < 0 {
				return nil, fmt.Errorf("%v<- incorrect brackets balance, could not close not opened bracket", input[0:i+1])
			}

			if counter == 0 {
				argument, err := BracketsToExpressionTree(input[from+1 : i])
				if err != nil {
					return nil, err
				}
				if argument != nil {
					root.Arguments = append(root.Arguments, argument)
				}
			}

		}

		if input[i] != '(' && input[i] != ')' {
			return nil, fmt.Errorf("%v<- unknown symbol", input[0:i+1])
		}
	}

	if counter != 0 {
		return nil, fmt.Errorf("number of opened brackets are not equal to closed (difference=%v)", counter)
	}

	return &root, nil
}

func main() {}
