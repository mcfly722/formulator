package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/*
func Test_Base3(t *testing.T) {
	for i := int64(0); i < 100; i++ {
		fmt.Println(fmt.Sprintf("%3v: %v", i, Base23(i)))
	}
}
*/
func Test_Number2CalculationTree(t *testing.T) {
	for i := int64(0); i < 10; i++ {
		tree := Number2CalculationTree(i)
		root, _ := json.Marshal(&tree)
		fmt.Println(fmt.Sprintf("%3v: %v, %v\n\n\n", i, Base23(i), string(root)))
	}
}
