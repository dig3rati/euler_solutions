package problem_5

import (
	"common"
)

func SmallestEvenlyDivisibleNumber(numbers []uint64) uint64 {
	r := uint64(1)
	for _, n := range numbers {
		r = common.LCM(r, n)
	}
	return r
}
