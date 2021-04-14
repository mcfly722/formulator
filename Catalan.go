package main

import "fmt"

// BracketStep node of binary tree
type BracketStep struct {
	x int
	y int
}

// Node of binary tree
type Node struct {
	Parent *Node
	Childs []*Node
}

func newNode(parent *Node) *Node {
	node := &Node{
		Parent: parent,
		Childs: []*Node{}}
	return node
}

// ToString converts binary tree with root to string representation
func (node *Node) ToString() string {
	if len(node.Childs) == 0 {
		return "*"
	}

	output := ""
	for _, n := range node.Childs {
		output += n.ToString()
	}
	return fmt.Sprintf("(%v)", output)
}

func printBrackets(tail []BracketStep) {
	treeRoot := BracketsStepsToBinaryTree(tail)

	root, err := StringToBinaryTree(treeRoot.ToString(), nil)

	if err != nil {
		fmt.Println(fmt.Sprintf("%v -> %v error:%v", BracketsStepsToString(tail), treeRoot.ToString(), err))
	}

	nextBrackets, err := getNextBracketsTree(BracketsStepsToString(tail))
	if err != nil {
		fmt.Println(fmt.Sprintf("%v error:%v", BracketsStepsToString(tail), err))
	}

	fmt.Println(fmt.Sprintf("%v -> %v -> %v equal=%v %v", BracketsStepsToString(tail), treeRoot.ToString(), root.ToString(), treeRoot.ToString() == root.ToString(), nextBrackets))
}

// BracketsStepsToBinaryTree converts brackets steps to binary tree
func BracketsStepsToBinaryTree(tail []BracketStep) *Node {
	root := newNode(nil)
	openedBracketSequence := []*Node{root}

	for _, step := range tail {

		for i := 0; i < step.x; i++ {
			current := openedBracketSequence[len(openedBracketSequence)-1]
			node := newNode(current)
			current.Childs = append(current.Childs, node)
			openedBracketSequence = append(openedBracketSequence, node)
		}

		for i := 0; i < step.y; i++ {
			validParent := openedBracketSequence[len(openedBracketSequence)-1].Parent
			for ; len(validParent.Childs) == 2; validParent = validParent.Parent {
			} // search last parent with one operand

			node := newNode(validParent)
			validParent.Childs = append(validParent.Childs, node)
			openedBracketSequence = append(openedBracketSequence, node)
		}
	}
	return root
}

// StringToBinaryTree deserialize from string to binary tree
func StringToBinaryTree(input string, parent *Node) (*Node, error) {
	node := newNode(parent)

	if input == "*" {
		return node, nil
	}

	if input[0] == '(' && input[len(input)-1] == ')' {
		pos := 1
		bracketsCounter := 1
		for ; bracketsCounter > 0 && pos < len(input); pos++ {
			if input[pos] == '(' {
				bracketsCounter++
			}
			if input[pos] == ')' {
				bracketsCounter--
			}
		}
		if bracketsCounter > 0 {
			return nil, fmt.Errorf("Opened brackets are not equal to closed in '%v' diff=%v", input, bracketsCounter)
		}

		if pos == len(input) {
			return StringToBinaryTree(input[1:len(input)-1], node)
		}

		childNode1, err := StringToBinaryTree(input[0:pos], node)
		if err != nil {
			return nil, err
		}

		childNode2, err := StringToBinaryTree(input[pos:len(input)], node)
		if err != nil {
			return nil, err
		}

		node.Childs = append(node.Childs, childNode1)
		node.Childs = append(node.Childs, childNode2)
		return node, nil
	}

	if input[0] == '*' {
		node.Childs = append(node.Childs, newNode(parent))
		childNode, err := StringToBinaryTree(input[1:len(input)], node)
		if err != nil {
			return nil, err
		}
		node.Childs = append(node.Childs, childNode)
		return node, nil
	}

	if input[0] == '(' && input[len(input)-1] == '*' {
		childNode, err := StringToBinaryTree(input[1:len(input)-2], node)
		if err != nil {
			return nil, err
		}
		node.Childs = append(node.Childs, childNode)
		node.Childs = append(node.Childs, newNode(parent))
		return node, nil
	}

	return nil, fmt.Errorf("Could not parse '%v'", input)
}

// StringToBracketsSteps converts brackets to steps
func StringToBracketsSteps(input string) ([]BracketStep, int, error) {
	steps := []BracketStep{}
	i, x, y := 0, 0, 0
	for i < len(input) {
		step := BracketStep{
			x: 0, y: 0,
		}

		for ; i < len(input) && input[i] == '('; i++ {
			step.x++
			x++
		}

		for ; i < len(input) && input[i] == ')'; i++ {
			step.y++
			y++
		}

		if i < len(input) && input[i] != '(' && input[i] != ')' {
			return nil, -1, fmt.Errorf("%v<- unknown symbol", input[0:i])
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
	startY := 1

	if !previousSolutionAlreadyReached && len(srcTail) > 0 {
		startX = srcTail[0].x
		startY = srcTail[0].y
	}

	for i := startX; i <= sizeX; i++ {
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

func getNextBracketsTree(input string) (string, error) {
	tail, size, err := StringToBracketsSteps(input)
	if err != nil {
		return "", err
	}

	nextTail, _, solutionFound := getNextStepsTail(tail, []BracketStep{}, size, size, false)

	if !solutionFound {
		return fmt.Sprintf("%v -> %v", tail, "next solution not found"), nil
	}

	return fmt.Sprintf("%v -> %v", tail, nextTail), nil
}

// BracketsStepsToString represent bracket steps as string
func BracketsStepsToString(tail []BracketStep) string {
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

// NextBracket recursively generates bracket ledder (each element is one ledder step with width and height)
func NextBracket(tail []BracketStep, sizeX int, sizeY int, output func([]BracketStep)) {
	//fmt.Println(fmt.Sprintf("next bracket"))

	if sizeX == 0 && sizeY == 0 {
		output(tail)
		return
	}

	for i := 1; i <= sizeX; i++ {
		for j := 1; j <= sizeY-(sizeX-i); j++ {

			//fmt.Println(fmt.Sprintf("%v,%v, %v,%v", i, j, size_x, size_y))
			nextTail := append(tail, BracketStep{x: i, y: j})

			NextBracket(nextTail, sizeX-i, sizeY-j, output)
		}
	}
}

func main() {
	NextBracket([]BracketStep{}, 4, 4, printBrackets)
}
