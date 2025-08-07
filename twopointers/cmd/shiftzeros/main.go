package main

import (
	"fmt"

	"github.com/shotokan/coding-interview-patterns/twopointers/shiftzeros"
)

func main() {
	// values := []int{0, 1, 0, 3, 2}
	values := []int{1, 4, 0, 5, 0, 3, 2}
	result := shiftzeros.ShiftZerosToEnd(values)
	fmt.Println(result)
}
