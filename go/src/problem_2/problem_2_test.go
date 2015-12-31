package problem_2

import (
	"fmt"
	"testing"
)

func TestProblem2(t *testing.T) {
	sum := SumOfFibonacciSeq(func(n uint64) bool {
		return n%2 == 0
	}, 4000000)
	fmt.Printf("Sum of even fibinacci sequence under 4 million is %v\n", sum)
}
