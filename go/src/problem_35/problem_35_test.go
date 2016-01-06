package problem_35

import (
	"testing"
)

func TestProblem35(t *testing.T) {
	checks := []struct {
		below  uint64
		expect uint64
	}{
		{100, 13},
		{1000000, 55},
	}
	for _, c := range checks {
		num := NumberOfCircularPrimes(c.below)
		t.Logf("Number of circular primes below %+v are %+v\n", c.below, num)
		if c.expect != num {
			t.Errorf("Expectation failed:\nExpected: %+v\nGot: %+v\n", c.expect, num)
		}
	}
}
