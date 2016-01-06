package problem_20

import (
	"math/big"
	"testing"
)

func TestProblem20(t *testing.T) {
	checks := []struct {
		n      int64
		expect uint64
	}{
		{10, 27},
		{100, 648},
	}
	for _, c := range checks {
		dsum := FactorialDigitSum(big.NewInt(c.n))
		t.Logf("Sum of the digits in the number %+v! is %+v\n", c.n, dsum)
		if c.expect != dsum {
			t.Errorf("Expectation failed:\nExpected: %+v\nGot: %+v\n", c.expect, dsum)
		}
	}
}
