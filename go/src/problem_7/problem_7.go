package problem_7

import (
	"common"
)

func NthPrime(n uint64) uint64 {
	_, below := common.RangeOfNthPrime(n)
	if n < 10 {
		below = n * n
	}
	out := common.PrimeSieveOfEratV1(below)
	var nth_prime, count uint64
	for p := range out {
		count++
		if count == n {
			nth_prime = p
		}
	}
	return nth_prime
}
