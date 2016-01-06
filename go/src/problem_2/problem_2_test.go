package problem_2

import (
	"testing"
)

func TestProblem2(t *testing.T) {
	checks := []struct {
		condition func(uint64) bool
		upto      uint64
		expect    uint64
	}{
		{func(n uint64) bool { return n%2 == 0 }, 4000000, 4613732},
	}
	for _, c := range checks {
		sum := SumOfFibonacciSeq(c.condition, c.upto)
		t.Logf("Sum of even fibinacci sequence under %v is %v\n", c.upto, sum)
		if c.expect != sum {
			t.Errorf("Expectation failed:\nExpected: %+v\nGot: %+v\n", c.expect, sum)
		}
	}
}
