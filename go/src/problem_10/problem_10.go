package problem_10

import (
	"common"
)

func SumOfPrimes(below uint64) uint64 {
	out := common.PrimeSeq(below)
	var sum uint64
	for n := range out {
		sum += n
	}
	return sum
}
