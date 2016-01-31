package problem_25

import (
	"testing"
)

func TestProblem25(t *testing.T) {
	checks := []struct {
		digits int
		expect int
	}{
		{2, 7},
		{3, 12},
		{1000, 4782},
	}
	for _, c := range checks {
		index := FirstFibonacciNumberWithDigits(c.digits)
		t.Logf("Index of first fibonacci number with %d digits is %+v\n", c.digits, index)
		if c.expect != index {
			t.Errorf("Expectation failed:\nExpected: %+v\nGot: %+v\n", c.expect, index)
		}
	}
}
