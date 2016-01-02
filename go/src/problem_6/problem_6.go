package problem_6

import (
	"common"
)

func DiffOfSquares(n uint64) uint64 {
	a := common.SumOfNaturalNumbers(n)
	b := common.SumOfNaturalSquares(n)
	return a*a - b
}
