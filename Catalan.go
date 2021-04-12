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

	fmt.Println(fmt.Sprintf("%v -> %v -> %v equal=%v", BracketsStepsToString(tail), treeRoot.ToString(), root.ToString(), treeRoot.ToString() == root.ToString()))

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
