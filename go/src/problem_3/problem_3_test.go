package problem_3

import (
	"fmt"
	"testing"
)

func TestProblem3(t *testing.T) {
	fac := LargestPrimeFactor(600851475143)
	fmt.Printf("Largest prime factor of 600851475143 is %v\n", fac)
}
