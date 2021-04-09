package main

import "fmt"

// Bracket node of binary tree
type BracketStep struct {
	x int
	y int
}

func printBrackets(tail []BracketStep) {
	fmt.Println(BracketsStepsToString(tail))
}

// BracketsStepsToString represent bracket steps as string
func BracketsStepsToString(tail []BracketStep) string {
	output:=""
	for _,step := range tail {
		for i:=0;i<step.x;i++{
			output+=fmt.Sprintf("(");
		}
		for i:=0;i<step.y;i++{
			output+=fmt.Sprintf(")");
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
