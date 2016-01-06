package problem_6

import (
	"testing"
)

func TestProblem6(t *testing.T) {
	checks := []struct {
		n      uint64
		expect uint64
	}{
		{10, 2640},
		{100, 25164150},
	}
	for _, c := range checks {
		diff := DiffOfSquares(c.n)
		t.Logf("Difference between the sum of the squares of the \n"+
			"first %+v natural numbers and the square of the sum is %+v\n", c.n, diff)
		if c.expect != diff {
			t.Errorf("Expectation failed:\nExpected: %+v\nGot: %+v\n", c.expect, diff)
		}
	}
}
