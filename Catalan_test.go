package main

import (
	"testing"
)

func Test_Brackets(t *testing.T) {
	NextBracket([]BracketStep{}, 6, 6, printBrackets)
}
