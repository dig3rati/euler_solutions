package problem_1

import (
	"testing"
)

func TestProblem1(t *testing.T) {
	checks := []struct {
		multiples []uint64
		upto      uint64
		expect    uint64
	}{
		{[]uint64{3, 5}, 10, 23},
		{[]uint64{3, 5}, 1000, 233168},
	}
	for _, c := range checks {
		sum := SumOfMultiples(c.multiples, c.upto)
		t.Logf("Sum of multiples of %+v below %d is %v\n", c.multiples, c.upto, sum)
		if c.expect != sum {
			t.Errorf("Expectation failed:\nExpected: %+v\nGot: %+v\n", c.expect, sum)
		}
	}
}
