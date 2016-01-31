package problem_25

import (
	"common"
)

func FirstFibonacciNumberWithDigits(digits int) int {
	fibseq := common.FibonacciContSeq()
	i := int(1)
	for n := range fibseq {
		if common.NumberOfDigits(n) == digits {
			break
		} else {
			i++
		}
	}
	return i
}
