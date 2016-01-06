package problem_50

import (
	"testing"
)

func TestConsecutivePrimeSum(t *testing.T) {
	checks := []struct {
		below  uint64
		expect uint64
	}{
		{100, 41},
		{1000, 953},
		{1000000, 997651},
	}
	for _, c := range checks {
		sum := ConsecutivePrimeSum(c.below)
		t.Logf("Prime number %+v, below %+v, can be written as the sum of the most consecutive primes", sum, c.below)
		if c.expect != sum {
			t.Errorf("Expectation failed:\nExpected: %+v\nGot: %+v\n", c.expect, sum)
		}
	}
}
