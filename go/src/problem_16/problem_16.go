package problem_16

import (
	"common"
	"math/big"
	"strconv"
)

func PowerDigitSum(x, y *big.Int) uint64 {
	pow := common.Pow(x, y)
	pbytes := []byte(pow.String())
	sum := uint64(0)
	for _, b := range pbytes {
		n, _ := strconv.ParseUint(string(b), 10, 64)
		sum += n
	}
	return sum
}
