package zeroOneTwoTree

// Node of tree
type Node struct {
	Parent *Node
	Childs []*Node
}

// Point represent number of opened and closed brackets
type Point struct {
	Opens  int
	Closes int
}

func recursion(_x int, _y int, bracketsStack []Point, diagonal [32]int, maxBracketPairs int, ready func(bracketsStack []Point, diagonal [32]int)) {

	if _x == maxBracketPairs && _y == maxBracketPairs {
		//ready
		ready(bracketsStack, diagonal)
	} else {
		if diagonal[_x-_y] < 2 {
			for x := _x + 1; x <= maxBracketPairs; x++ {
				diagonal[x-_y-1] = diagonal[x-_y-1] + 1
				for y := _y + 1; y <= maxBracketPairs; y++ {
					if y <= x {
						if diagonal[x-y] < 3 {
							newBracketsStack := append(bracketsStack, Point{Opens: x - _x, Closes: y - _y})
							recursion(x, y, newBracketsStack, diagonal, maxBracketPairs, ready)
						}
					}
				}
			}

		}

	}

}

// Recombine all binary trees
func Recombine(maxBracketPairs int, ready func(bracketsStack []Point, diagonal [32]int)) {
	diagonal := [32]int{}
	recursion(0, 0, []Point{}, diagonal, maxBracketPairs, ready)
}
