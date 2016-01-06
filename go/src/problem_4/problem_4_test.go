package problem_4

import (
	"testing"
)

func TestProblem4(t *testing.T) {
	checks := []struct {
		digits uint64
		expect uint64
	}{
		{2, 9009},
		{3, 906609},
	}
	for _, c := range checks {
		n := LargestPalindrome(c.digits)
		t.Logf("Largest palindrome from the product of two %d-digit numbers is %+v\n",
			c.digits, n)
		if c.expect != n {
			t.Errorf("Expectation failed:\nExpected: %+v\nGot: %+v\n", c.expect, n)
		}
	}
}
