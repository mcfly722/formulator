package main

import (
	"fmt"

	"github.com/mcfly722/formulator/catalan"
)

func main() {

	brackets := "(()())((()))"

	for i := 0; i < 100000000; i++ {

		newBrackets, err := catalan.GetNextBracketsSequence(brackets)
		if err != nil {
			fmt.Println(fmt.Sprintf("error for %v : %v", brackets, err))
			break
		}
		fmt.Println(fmt.Sprintf("%5v  :   %v", i, newBrackets))
		brackets = newBrackets
	}

}
