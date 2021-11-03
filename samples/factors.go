package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
)

// Point ...
type Point struct {
	X int
	Y int
}

func main() {

	points := []Point{}

	for n := 1; n < 100; n++ {

		i := 0
		j := 0

		for ; !big.NewInt(int64(i)).ProbablyPrime(0); i = rand.Intn(100000) {
		}
		for ; !big.NewInt(int64(j)).ProbablyPrime(0); j = rand.Intn(100000) {
		}

		points = append(points, Point{X: i * j, Y: (i + j) / 2})

	}

	json, _ := json.MarshalIndent(points, "", "\t")
	fmt.Println(string(json))
}
