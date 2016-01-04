package problem_10

import (
	"testing"
)

func TestProblem10(t *testing.T) {
	sum := SumOfPrimes(2000000)
	t.Logf("Sum of all the primes below two million is %+v\n", sum)
}
