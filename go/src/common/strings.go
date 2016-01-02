package common

import (
	"fmt"
	"math"
)

func ReverseString(input string) string {
	s := []rune(input)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func IsPalindrome(input string) bool {
	return input == ReverseString(input)
}

func PalindromeSeq(digits uint64) chan uint64 {
	out := make(chan uint64)
	go func() {
		defer close(out)
		l := uint64(math.Pow(10, float64(digits)) - 1)
		for i := l; i > 0; i-- {
			for j := l; j > 0; j-- {
				n := j * i
				if IsPalindrome(fmt.Sprintf("%d", n)) {
					out <- n
				}
			}
		}
	}()
	return out
}
