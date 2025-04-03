package main

import (
	"fmt"
	"golang-course/lesson02"
)

func main() {
	fmt.Println(lesson02.FibonacciIterative(47))
	fmt.Println(lesson02.IsPrime(7))
	fmt.Println(lesson02.Increment("0000000"))
	fmt.Println(lesson02.IsBinaryPalindrome(3))
	fmt.Println(lesson02.FibonacciRecursive(47))
}
