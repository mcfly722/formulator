package main

import (
	"testing"
)

func Test_Brackets(t *testing.T) {
	NextBracket([]Node{}, 4, 4, printBrackets)
}
