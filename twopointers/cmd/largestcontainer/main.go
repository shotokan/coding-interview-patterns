package main

import (
	"fmt"

	"github.com/shotokan/coding-interview-patterns/twopointers/largestcontainer"
)

func main() {
	container := []int{2, 7, 8, 3, 7, 6}
	// container := []int{}
	// container := []int{3, 4, 3, 3, 3}

	result := largestcontainer.LargestContainer(container)
	fmt.Println(result)
}
