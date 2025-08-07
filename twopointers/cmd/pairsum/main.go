package main

import (
	"fmt"

	"github.com/shotokan/coding-interview-patterns/twopointers/pairsum"
)

func main() {
	// nums := []int{-5, -2, 3, 4, 6}
	nums := []int{1, 1, 1}
	resp := pairsum.FindPairSum(nums, 3)
	fmt.Println(resp)
}
