package problem_6

import (
	"fmt"
	"testing"
)

func TestProblem6(t *testing.T) {
	d := DiffOfSquares(100)
	fmt.Printf("Difference between the sum of the squares of the \n"+
		"first one hundred natural numbers and the square of the sum is %+v\n", d)
}
