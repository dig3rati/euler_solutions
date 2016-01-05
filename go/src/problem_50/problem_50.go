package problem_50

import (
	"common"
)

func ConsecutivePrimeSum(below uint64) uint64 {
	out := common.PrimeSieveOfEratV1(below)
	var max, sum uint64
	for p := range out {
		sum += p
	}
	return max
}
