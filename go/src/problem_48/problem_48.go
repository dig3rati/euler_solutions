package problem_48

import (
	"common"
	"math/big"
)

func SumOfSelfPowers(n *big.Int) *big.Int {
	return common.SumOfSelfPowers(big.NewInt(1), n)
}
