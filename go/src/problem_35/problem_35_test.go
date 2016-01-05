package problem_35

import (
	"testing"
)

func TestProblem35(t *testing.T) {
	t.Logf("Number of circular primes below one million are %d\n",
		NumberOfCircularPrimes(1000000))
}
