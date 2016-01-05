package problem_35

import (
	"common"
)

func NumberOfCircularPrimes(below uint64) uint64 {
	out := common.PrimeSieveOfEratV1(below)
	count := uint64(0)
	for p := range out {
		if common.IsCircularPrime(p) {
			count++
		}
	}
	return count
}
