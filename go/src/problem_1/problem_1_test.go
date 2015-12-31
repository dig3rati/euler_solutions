package problem_1

import (
	"fmt"
	"testing"
)

func TestProblem1(t *testing.T) {
	sum := SumOfMultiples([]uint64{3, 5}, 1000)
	fmt.Printf("Sum of multiples of 3, 5 below 1000 is %v\n", sum)
}
