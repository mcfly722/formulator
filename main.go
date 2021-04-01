package main

import (
	"encoding/json"
	"fmt"
	"math"
)

func main() {}

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
	First    *Operation
	Second   *Operation
}

func newOperation(n int) *Operation {
	return &Operation{N: n}
}

func (root *Operation) newChilds(n int, base byte, layer *[]*Operation) int {
	if base == 1 {
		root.First = newOperation(n + 1)
		(*layer) = append(*layer, root.First)

	}

	if base == 2 {
		root.First = newOperation(n + 1)
		root.Second = newOperation(n + 2)
		(*layer) = append(*layer, root.First ,root.Second)
	}

	return n+1
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
	//fmt.Println(fmt.Sprintf("currentLayerLen:%v", len(currentLayer)))
	//fmt.Println(fmt.Sprintf("LenBase:%v i:%v", len(base),i))


	for i < len(base) && len(currentLayer) > 0 {
		//fmt.Println(fmt.Sprintf("i:%v", i))

		nextLayer := []*Operation{}

		for j := 0; j < len(currentLayer); j++ {

			if i < len(base) {
				i = currentLayer[j].newChilds(i, base[i], &nextLayer)
			} else {
				i = currentLayer[j].newChilds(i, 0, &nextLayer)
			}

			i++
		}

		//fmt.Println(fmt.Sprintf("CurrentLayer:%v", currentLayer))

		currentLayer=nextLayer

	}
	return root
}
