package problem_50

import (
	"testing"
)

func TestConsecutivePrimeSum(t *testing.T) {
	t.Logf("Prime number %+v, below one-million, can be written as the sum of the most consecutive primes", ConsecutivePrimeSum(1000000))
}
