package problem_3

import (
	"common"
)

func LargestPrimeFactor(number uint64) uint64 {
	factors := common.PrimeFactors(number)
	var largest uint64
	for f := range factors {
		if largest < f {
			largest = f
		}
	}
	return largest
}
