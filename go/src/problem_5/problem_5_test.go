package problem_5

import (
	"testing"
)

func TestProblem5(t *testing.T) {
	n := SmallestEvenlyDivisibleNumber([]uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	t.Logf("Smallest positive number that is evenly divisible by 1 to 20 is %+v\n", n)
}
