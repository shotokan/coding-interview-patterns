package main

import (
	"fmt"

	"github.com/shotokan/coding-interview-patterns/twopointers/tripletsum"
)

func main() {
	// nums := []int{-4, -4, -2, 0, 0, 1, 2, 3}
	// nums := []int{1, 1, 1}
	nums := []int{0, 0, 1, -1, 1, -1}
	// nums := []int{0, 0, 0}
	resp := tripletsum.FindTripletSum(nums)
	fmt.Println(resp)
}
