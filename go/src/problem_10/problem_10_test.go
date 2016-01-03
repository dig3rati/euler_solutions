package problem_10

import (
	"fmt"
	"testing"
)

func TestProblem10(t *testing.T) {
	sum := SumOfPrimes(2000000)
	fmt.Printf("Sum of all the primes below two million is %+v\n", sum)
}
