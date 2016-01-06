package problem_5

import (
	"testing"
)

func TestProblem5(t *testing.T) {
	checks := []struct {
		numbers []uint64
		expect  uint64
	}{
		{[]uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2520},
		{[]uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 232792560},
	}
	for _, c := range checks {
		n := SmallestEvenlyDivisibleNumber(c.numbers)
		t.Logf("Smallest positive number that is evenly divisible by %+v is %+v\n", c.numbers, n)
		if c.expect != n {
			t.Errorf("Expectation failed:\nExpected: %+v\nGot: %+v\n", c.expect, n)
		}
	}
}
