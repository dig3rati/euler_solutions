package common

import (
	"math"
	"math/big"
	"testing"
)

func TestFibonacciSeq(t *testing.T) {
	cases := []struct {
		start, second, upto uint64
		want                []uint64
	}{
		{1, 2, 10, []uint64{1, 2, 3, 5, 8}},
		{10, 20, 50, []uint64{10, 20, 30, 50}},
	}
	for _, c := range cases {
		out := FibonacciSeq(c.start, c.second, c.upto)
		for _, w := range c.want {
			if w != <-out {
				t.Errorf("Failed to check case: %+v\n", c)
			}
		}
	}
}

func TestIsPrime(t *testing.T) {
	cases := []struct {
		number uint64
		want   bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{10, false},
		{math.MaxUint32, false},
	}
	for _, c := range cases {
		if IsPrime(c.number) != c.want {
			t.Errorf("Failed to check case: %+v\n", c)
		}
	}
}

func TestRangeOfNthPrime(t *testing.T) {
	cases := []struct {
		n    uint64
		want struct{ above, below uint64 }
	}{
		{n: 10, want: struct{ above, below uint64 }{20, 30}},
	}
	for _, c := range cases {
		a, b := RangeOfNthPrime(c.n)
		if c.want.above != a || c.want.below != b {
			t.Errorf("Failed to check case: %+v\n", c)
		}
	}
}

func TestPrimeSieveOfEratV1(t *testing.T) {
	cases := []struct {
		upto uint64
		want []uint64
	}{
		{7, []uint64{2, 3, 5, 7}},
		{10, []uint64{2, 3, 5, 7}},
	}
	for _, c := range cases {
		out := PrimeSieveOfEratV1(c.upto)
		for _, w := range c.want {
			if w != <-out {
				t.Errorf("Failed to check case: %+v\n", c)
			}
		}
	}
}

func TestPrimeFactors(t *testing.T) {
	cases := []struct {
		number uint64
		want   []uint64
	}{
		{13195, []uint64{5, 7, 13, 29}},
	}
	for _, c := range cases {
		out := PrimeFactors(c.number)
		for _, w := range c.want {
			if w != <-out {
				t.Errorf("Failed to check case: %+v\n", c)
			}
		}
	}
}

func TestGCD(t *testing.T) {
	cases := []struct {
		a, b uint64
		want uint64
	}{
		{2, 3, 1},
		{7, 21, 7},
	}
	for _, c := range cases {
		if GCD(c.a, c.b) != c.want {
			t.Errorf("Failed to check case: %+v\n", c)
		}
	}
}

func TestLCM(t *testing.T) {
	cases := []struct {
		a, b uint64
		want uint64
	}{
		{2, 3, 6},
		{65, 15, 195},
		{3, 21, 21},
	}
	for _, c := range cases {
		if LCM(c.a, c.b) != c.want {
			t.Errorf("Failed to check case: %+v\n", c)
		}
	}
}

func TestFactorial(t *testing.T) {
	cases := []struct {
		n, want *big.Int
	}{
		{big.NewInt(10), big.NewInt(3628800)},
	}
	for _, c := range cases {
		if c.want.Cmp(Factorial(c.n)) != 0 {
			t.Errorf("Failed to check case: %+v\n", c)
		}
	}
}
