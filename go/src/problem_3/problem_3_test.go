package problem_3

import (
	"testing"
)

func TestProblem3(t *testing.T) {
	fac := LargestPrimeFactor(600851475143)
	t.Logf("Largest prime factor of 600851475143 is %v\n", fac)
}
