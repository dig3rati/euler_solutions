package problem_4

import (
	"testing"
)

func TestProblem4(t *testing.T) {
	n := LargestPalindrome(3)
	t.Logf("Largest palindrome from the product of two 3-digit numbers is %+v\n", n)
}
