package problem_3

import (
	"testing"
)

func TestProblem3(t *testing.T) {
	checks := []struct {
		number uint64
		expect uint64
	}{
		{13195, 29},
		{600851475143, 6857},
	}
	for _, c := range checks {
		fac := LargestPrimeFactor(c.number)
		t.Logf("Largest prime factor of %+v is %v\n", c.number, fac)
		if c.expect != fac {
			t.Errorf("Expectation failed:\n Expected: %+v\nGot: %+v\n", c.expect, fac)
		}
	}
}
