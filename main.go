package main

import (
	"encoding/json"
	"fmt"
	"math"
)

func main() {
	for i := int64(0); i < 5; i++ {
		tree := Number2CalculationTree(i)
		root, _ := json.Marshal(&tree)
		fmt.Println(fmt.Sprintf("%3v: %v, %v\n\n\n", i, Base23(i), string(root)))
	}
}

func reverseByteSlice(numbers []byte) []byte {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

// Base23 converts int64 number to array of bits with base3 without zeros on sides except first zero
func Base23(t int64) []byte {

	if t == 0 {
		return []byte{0}
	}

	t--

	result := []byte{}
	result = append(result, byte(t%2+1))

	i := 1
	p := int64(math.Floor((float64)(t / 2)))
	if p > 0 {
		for {
			result = append(result, byte(p%3))
			p = int64(math.Floor((float64)(p / 3)))
			i++
			if p == 0 {
				return reverseByteSlice(result)
			}
		}
	}
	return reverseByteSlice(result)
}

// Operation structure
type Operation struct {
	N        int
	Operands []*Operation
}

func newOperation(n int) *Operation {
	return &Operation{N: n}
}

func (root *Operation) newChilds(n int, base byte, layer *[]*Operation) int {
	if base == 1 {
		operation := newOperation(n + 1)
		root.Operands = append(root.Operands, operation)
		(*layer) = append(*layer, operation)

	}

	if base == 2 {
		operation1 := newOperation(n + 1)
		operation2 := newOperation(n + 2)
		root.Operands = append(root.Operands, operation1, operation2)
		(*layer) = append(*layer, operation1, operation2)
		n++
	}
	n = n + 1
	return n
}

func (root *Operation) toJSON(prefix string) string {
	j, _ := json.Marshal(&root)
	return fmt.Sprintf("%v%v", prefix, string(j))
}

// Number2CalculationTree generates calculation tree based on number
func Number2CalculationTree(n int64) *Operation {
	base := Base23(n)

	fmt.Println(fmt.Sprintf("%3v: %v.", n, base))

	root := newOperation(0)
	currentLayer := []*Operation{}
	i := root.newChilds(0, base[0], &currentLayer)

	for i < len(base) && len(currentLayer) > 0 {

		nextLayer := []*Operation{}

		for j := 0; j < len(currentLayer); j++ {

			if i < len(base) {
				i = currentLayer[j].newChilds(i, base[i], &nextLayer)
			} else {
				i = currentLayer[j].newChilds(i, 0, &nextLayer)
			}

			i++
		}

		copy(currentLayer, nextLayer)

	}
	return root
}
