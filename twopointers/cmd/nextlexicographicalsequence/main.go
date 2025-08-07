package main

import (
	"fmt"

	"github.com/shotokan/coding-interview-patterns/twopointers/nextlexicographicalsequence"
)

func main() {
	//input := "abcd"
	input := "dcba"

	next := nextlexicographicalsequence.NextLexicographicalSequence(input)
	fmt.Println(next)

}
