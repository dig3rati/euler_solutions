package problem_50

import (
	"common"
)

func ConsecutivePrimeSum(below uint64) uint64 {
	pchan := common.PrimeSieveOfEratV1(below)
	primes := make([]uint64, 0, common.NumberOfPrimes(below))
	for p := range pchan {
		primes = append(primes, p)
	}
	var sum, count, fsum, fcount uint64
	for i, _ := range primes {
		sum, count = 0, 0
		for _, cp := range primes[i:] {
			sum += cp
			if sum >= below {
				break
			}
			if common.IsPrime(sum) && count > fcount {
				fcount = count
				fsum = sum
			}
			count++
		}
	}
	return fsum
}
