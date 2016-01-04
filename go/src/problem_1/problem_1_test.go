package problem_1

import (
	"testing"
)

func TestProblem1(t *testing.T) {
	sum := SumOfMultiples([]uint64{3, 5}, 1000)
	t.Logf("Sum of multiples of 3, 5 below 1000 is %v\n", sum)
}
