package problem_20

import (
	"common"
	"fmt"
	"math/big"
)

func FactorialDigitSum(n *big.Int) uint64 {
	fac := common.Factorial(n)
	return common.SumOfDigits(fmt.Sprintf("%d", fac))
}
