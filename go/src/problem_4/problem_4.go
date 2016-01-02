package problem_4

import (
	"common"
)

func LargestPalindrome(digits uint64) uint64 {
	out := common.PalindromeSeq(digits)
	var largest uint64
	for p := range out {
		if p > largest {
			largest = p
		}
	}
	return largest
}
