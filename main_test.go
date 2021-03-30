package main

import (
	"fmt"
	"testing"
)

func Test_Base3(t *testing.T) {
	for i := int64(0); i < 100; i++ {
		fmt.Println(fmt.Sprintf("%3v: %v", i, Base23(i)))
	}
}
