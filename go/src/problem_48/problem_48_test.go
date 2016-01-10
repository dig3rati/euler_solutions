package problem_48

import (
	"math/big"
	"testing"
)

func TestProblem48(t *testing.T) {
	checks := []struct {
		n    *big.Int
		want *big.Int
	}{
		{big.NewInt(10), big.NewInt(10405071317)},
		{big.NewInt(1000), nil},
	}
	for _, c := range checks {
		sum := SumOfSelfPowers(c.n)
		t.Logf("Sum of self powers of 1 to %+v is %+v\n", c.n, sum)
		if c.want != nil && c.want.Cmp(sum) != 0 {
			t.Errorf("Expectation failed:\nExpected: %+v\nGot: %+v\n", c.want, sum)
		}
	}
}
