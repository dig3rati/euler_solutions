package problem_10

import (
	"testing"
)

func TestProblem10(t *testing.T) {
	checks := []struct {
		below  uint64
		expect uint64
	}{
		{10, 17},
		{2000000, 142913828922},
	}
	for _, c := range checks {
		sum := SumOfPrimes(c.below)
		t.Logf("Sum of all the primes below %+v is %+v\n", c.below, sum)
		if c.expect != sum {
			t.Errorf("Expectation failed:\nExpected: %+v\nGot: %+v\n", c.expect, sum)
		}
	}
}
