package lesson02

import (
	"strconv"
	"testing"
)

func BenchmarkIncrement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strBinary := strconv.FormatInt(int64(i), 2)
		Increment(strBinary)
	}
}

func BenchmarkIncrementWithStrConv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strBinary := strconv.FormatInt(int64(i), 2)
		IncrementWithStrConv(strBinary)
	}
}

func BenchmarkIsBinaryPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsBinaryPalindrome(i)
	}
}

func BenchmarkIsBinaryPalindromeWithoutBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsBinaryPalindromeWithoutBuilder(i)
	}
}
