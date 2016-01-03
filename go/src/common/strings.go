package common

import (
	"fmt"
	"math"
	"strconv"
)

// Reverses the input string
func ReverseString(input string) string {
	s := []rune(input)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

// Checks if the input string is a palindrome
func IsPalindrome(input string) bool {
	return input == ReverseString(input)
}

// Returns a channel for all palindromes for multiples of given digit numbers
func PalindromeSeq(digits uint64) chan uint64 {
	out := make(chan uint64)
	go func() {
		defer close(out)
		l := uint64(math.Pow(10, float64(digits)) - 1)
		m := l - uint64(99)
		for i := l; i > m; i-- {
			for j := l; j > m; j-- {
				n := j * i
				if IsPalindrome(fmt.Sprintf("%d", n)) {
					out <- n
				}
			}
		}
	}()
	return out
}

func SumOfDigits(n string) uint64 {
	bytes := []byte(n)
	var sum uint64
	for _, b := range bytes {
		v, _ := strconv.Atoi(string(b))
		sum += uint64(v)
	}
	return sum
}
