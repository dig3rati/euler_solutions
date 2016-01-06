package problem_7

import (
	"testing"
)

func TestProblem7(t *testing.T) {
	checks := []struct {
		n      uint64
		expect uint64
	}{
		{6, 13},
		{10001, 104743},
	}
	for _, c := range checks {
		p := NthPrime(c.n)
		t.Logf("%+v prime number is %+v\n", c.n, p)
		if c.expect != p {
			t.Errorf("Expectation failed:\nExpected: %+v\nGot: %+v\n", c.expect, p)
		}
	}
}
