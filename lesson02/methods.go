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
		return 0
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
		return 0
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

func IncrementWithStrConv(num string) int {
	intDecimal, _ := strconv.ParseInt(num, 2, 64)
	return int(intDecimal) + 1
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

func IsBinaryPalindromeWithoutBuilder(n int) bool {
	binaryStr := strconv.FormatInt(int64(n), 2)
	strLen := len(binaryStr)

	binaryStrRune := []rune(binaryStr)
	reversedStrRune := []rune(binaryStr)

	for i := 0; i < strLen/2; i++ {
		reversedStrRune[i] = binaryStrRune[strLen-i-1]
		reversedStrRune[strLen-i-1] = binaryStrRune[i]
	}

	return Increment(binaryStr) == Increment(string(reversedStrRune))
}

func ValidParentheses(s string) bool {
	var filteredStr string

	for i := 0; i < len(s); i++ {
		if strings.Index("(){}[]", string(s[i])) != -1 {
			filteredStr += string(s[i])
		}
	}

	filteredStrLen := len(filteredStr)

	if filteredStrLen%2 != 0 {
		return false
	}

	var strWithoutPairs string

	allowedPairs := "() {} []"

	for i := 0; i < filteredStrLen-1; i++ {
		tempStr := string(filteredStr[i]) + string(filteredStr[i+1])

		if strings.Index(allowedPairs, tempStr) == -1 {
			strWithoutPairs += string(filteredStr[i])
			if (i + 1) == filteredStrLen-1 {
				strWithoutPairs += string(filteredStr[i+1])
			}
		} else {
			i = i + 1
		}
	}

	strWithoutPairsLen := len(strWithoutPairs)

	for i := 0; i < strWithoutPairsLen/2; i++ {
		tempPair := string(strWithoutPairs[i]) + string(strWithoutPairs[strWithoutPairsLen-1-i])

		if strings.Index(allowedPairs, tempPair) == -1 {
			return false
		}
	}

	return true
}
