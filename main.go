package main

import (
	"math"
)

func reverseByteSlice(numbers []byte) []byte {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

// Base23 converts int64 number to array of bits with base3 without zeros on sides
func Base23(t int64) []byte {

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

func main() {}
