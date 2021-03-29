package main

import (
	"fmt"
	"math"
)

func base3(t int64) ([]byte) {
	result:=[]byte{}
	i:=0
	p:=t
	for {
		
		result =append(result, byte(p % 3))
		p = int64(math.Floor((float64)(p / 3)))
		i++
		if p==0 {
			return result
		}
	}
}

func main() {
	var i int64
	
	for i=0;i<100;i++ {
		base :=base3(i)
		fmt.Println(fmt.Sprintf("%3v: %v",i,base))
	}
}