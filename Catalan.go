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

	fmt.Println(fmt.Sprintf("%v -> %v", BracketsStepsToString(tail), treeRoot.ToString()))

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
	NextBracket([]BracketStep{}, 6, 6, printBrackets)
}
