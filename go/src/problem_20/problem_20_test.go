package problem_20

import (
	"fmt"
	"math/big"
	"testing"
)

func TestProblem20(t *testing.T) {
	fmt.Printf("Sum of the digits in the number 100! is %+v\n", FactorialDigitSum(big.NewInt(100)))
}
