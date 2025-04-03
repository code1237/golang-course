package lesson02

import (
	"math"
	"strconv"
	"strings"
)

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	dividers := 1

	for i := 2; i <= n; i++ {
		if n%i == 0 {
			if dividers++; dividers > 2 {
				return false
			}
		}
	}

	return true
}

func FibonacciIterative(n int) int {
	if n < 1 {
		panic("n must be greater than zero")
	}

	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	prev, current := 0, 1

	for i := 2; i < n; i++ {
		current = prev + current
		prev = current - prev
	}

	return current
}

func FibonacciRecursive(n int) int {
	if n < 1 {
		panic("n must be greater than zero")
	}

	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func Increment(num string) int {
	strLen := len(num)
	decimal := 0

	for i := 0; i < strLen; i++ {
		curMultiplier, _ := strconv.Atoi(string(num[i]))
		decimal += curMultiplier * int(math.Pow(2, float64(strLen-1-i)))
	}

	return decimal + 1
}

func IsBinaryPalindrome(n int) bool {
	binaryStr := strconv.FormatInt(int64(n), 2)
	strLen := len(binaryStr)

	var reversedStringBuilder strings.Builder

	for i := 0; i < strLen; i++ {
		reversedStringBuilder.WriteString(string(binaryStr[strLen-i-1]))
	}

	return Increment(binaryStr) == Increment(reversedStringBuilder.String())
}
