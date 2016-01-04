package problem_6

import (
	"testing"
)

func TestProblem6(t *testing.T) {
	d := DiffOfSquares(100)
	t.Logf("Difference between the sum of the squares of the \n"+
		"first one hundred natural numbers and the square of the sum is %+v\n", d)
}
