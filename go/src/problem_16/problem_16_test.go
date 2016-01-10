package problem_16

import (
	"math/big"
	"testing"
)

func TestProblem16(t *testing.T) {
	checks := []struct {
		x, y *big.Int
		want uint64
	}{
		{big.NewInt(2), big.NewInt(15), 26},
		{big.NewInt(2), big.NewInt(1000), 0},
	}
	for _, c := range checks {
		sum := PowerDigitSum(c.x, c.y)
		t.Logf("Sum of the digits of the number %v^%v is %+v\n", c.x, c.y, sum)
		if c.want != sum {
			t.Errorf("Expectation failed:\nExpected: %+v\nGot: %+v\n", c.want, sum)
		}
	}
}
