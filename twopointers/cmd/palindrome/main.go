package main

import (
	"fmt"

	"github.com/shotokan/coding-interview-patterns/twopointers/palindrome"
)

func main() {
	s := "!ab!a"
	isPalindrome := palindrome.IsPalindrome(s)
	fmt.Println(isPalindrome)
}
