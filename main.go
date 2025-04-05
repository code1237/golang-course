package main

import (
	"fmt"
	"golang-course/lesson02"
)

func main() {
	fmt.Println(lesson02.FibonacciIterative(47))
	fmt.Println(lesson02.IsPrime(7))
	fmt.Println(lesson02.Increment("11111"))
	fmt.Println(lesson02.IsBinaryPalindrome(3))
	fmt.Println(lesson02.FibonacciRecursive(10))
	fmt.Println(lesson02.ValidParentheses("([{}])"))
}
