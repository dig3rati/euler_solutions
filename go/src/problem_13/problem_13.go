package problem_13

import (
	"math/big"
)

func SumOfBigInts(nums []*big.Int) *big.Int {
	var sum *big.Int = big.NewInt(0)
	for _, i := range nums {
		sum.Add(sum, i)
	}
	return sum
}
