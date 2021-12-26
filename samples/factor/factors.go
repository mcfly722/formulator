package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
)

// Point ...
type Point struct {
	X float64
	Y float64
}

func main() {

	points := []Point{}

	for len(points) < 100 {

		x := rand.Intn(5000)
		y := rand.Intn(5000)

		if big.NewInt(int64(x)).ProbablyPrime(0) && big.NewInt(int64(y)).ProbablyPrime(0) {
			points = append(points, Point{X: float64(x * y), Y: float64((x + y) / 2)})
			//			fmt.Println(fmt.Sprintf("%6v * %6v -> %12v %12v", x, y, x*y, x+y))
		}

	}

	/*
		for i := 23; i < 100; i++ {
			for j := 23; j < 100; j++ {

				if big.NewInt(int64(i)).ProbablyPrime(0) && big.NewInt(int64(j)).ProbablyPrime(0) {
					points = append(points, Point{X: i * j, Y: (i + j) / 2})
				}
			}
		}
	*/
	json, _ := json.MarshalIndent(points, "", "\t")
	fmt.Println(string(json))
}
