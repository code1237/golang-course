package main

import (
	"fmt"
	"golang-course/lesson02"
)

func main() {
	fmt.Println(lesson02.FibonacciIterative(10))
	fmt.Println(lesson02.IsPrime(9))
	fmt.Println(lesson02.Increment("11111"))
	fmt.Println(lesson02.IncrementWithStrConv("11111"))
	fmt.Println(lesson02.IsBinaryPalindrome(7))
	fmt.Println(lesson02.IsBinaryPalindromeWithoutBuilder(7))
	fmt.Println(lesson02.FibonacciRecursive(10))
	fmt.Println(lesson02.ValidParentheses("func() { return fmt.Println(len([]int{1,2,3}))}"))
	fmt.Println(lesson02.ValidParentheses("func() { return fmt.Println(len)([]int{1,2,3}))}"))
}
