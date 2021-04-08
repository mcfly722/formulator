package main

import "fmt"

// Node node of binary tree
type Node struct {
	x int
	y int
}

func printBrackets(tail []Node) {
	fmt.Println(fmt.Sprintf("%v", tail))
}

// NextBracket recursively generates bracket ledder (each element is one ledder step with width and height)
func NextBracket(tail []Node, sizeX int, sizeY int, output func([]Node)) {
	//fmt.Println(fmt.Sprintf("next bracket"))

	if sizeX == 0 && sizeY == 0 {
		output(tail)
		return
	}

	for i := sizeX; i >= 1; i-- {
		for j := 1; j <= sizeY-(sizeX-i); j++ {

			//fmt.Println(fmt.Sprintf("%v,%v, %v,%v", i, j, size_x, size_y))
			nextTail := append(tail, Node{x: i, y: j})

			NextBracket(nextTail, sizeX-i, sizeY-j, output)
		}
	}
}

func main() {
	NextBracket([]Node{}, 4, 4, printBrackets)
}
