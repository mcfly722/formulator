package zeroOneTwoTree

import (
	"fmt"
)

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

func recursion(_x int, _y int, bracketsStack []Point, maxChilds int, diagonal [32]int, maxBracketPairs int, ready func(bracketsStack []Point, diagonal [32]int)) {

	if _x == maxBracketPairs && _y == maxBracketPairs {
		//ready
		ready(bracketsStack, diagonal)
	} else {
		if diagonal[_x-_y] < maxChilds {
			for x := _x + 1; x <= maxBracketPairs; x++ {
				diagonal[x-_y-1] = diagonal[x-_y-1] + 1
				for y := _y + 1; y <= maxBracketPairs; y++ {
					if y <= x {
						if diagonal[x-y] < maxChilds+1 {
							newBracketsStack := append(bracketsStack, Point{Opens: x - _x, Closes: y - _y})
							recursion(x, y, newBracketsStack, maxChilds, diagonal, maxBracketPairs, ready)
						}
					}
				}
			}

		}

	}

}

// Recombine all binary trees
func Recombine(maxBracketPairs int, maxChilds int, ready func(bracketsStack []Point, diagonal [32]int)) {
	diagonal := [32]int{}
	recursion(0, 0, []Point{}, maxChilds, diagonal, maxBracketPairs, ready)
}

func brackets2points(brackets string) (*[]Point, error) {
	totalOpens := 0
	totalCloses := 0

	points := []Point{}

	for i := 0; i < len(brackets); {

		opens := 0
		for ; i+opens < len(brackets) && brackets[i+opens] == '('; opens++ {
		}
		i = i + opens
		totalOpens = totalOpens + opens

		if i < len(brackets) && brackets[i] != ')' {
			return nil, fmt.Errorf("Unexpected Symbol %v <- Expecting '(' or ')'", brackets[:i+1])
		}

		closes := 0
		for ; i+closes < len(brackets) && brackets[i+closes] == ')'; closes++ {
		}
		i = i + closes
		totalCloses = totalCloses + closes

		if i < len(brackets) && brackets[i] != '(' {
			return nil, fmt.Errorf("Unexpected Symbol %v <- Expecting '(' or ')'", brackets[:i+1])
		}

		points = append(points, Point{Opens: opens, Closes: closes})

		if totalOpens < totalCloses {
			return nil, fmt.Errorf("%v <- total closes=%v are greater than opens=%v", brackets[:i], totalCloses, totalOpens)
		}
	}

	if totalOpens != totalCloses {
		return nil, fmt.Errorf("opened brackets=%v closed brackets=%v should be equal", totalOpens, totalCloses)
	}

	return &points, nil
}

// GetNextTree get current brackets representation of tree and return next one tree in brackets representation
func GetNextTree(brackets string) (string, error) {
	points, err := brackets2points(brackets)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", points), nil
}
