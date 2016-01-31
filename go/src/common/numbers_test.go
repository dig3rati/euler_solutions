package common

import (
	"math"
	"math/big"
	"reflect"
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

func TestFibonacciContSeq(t *testing.T) {
	cases := []struct {
		want []*big.Int
	}{
		{[]*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(2),
			big.NewInt(3), big.NewInt(5), big.NewInt(8)}},
	}
	for _, c := range cases {
		out := FibonacciContSeq()
		for _, n := range c.want {
			e := <-out
			if n.Cmp(e) != 0 {
				t.Errorf("Failed to check case: %+v\n", c)
			}
		}
	}
}

func TestCircularNumbers(t *testing.T) {
	cases := []struct {
		n    uint64
		want []uint64
	}{
		{123, []uint64{123, 231, 312}},
		{10, []uint64{10, 1}},
		{1, []uint64{1}},
	}
	for _, c := range cases {
		if !reflect.DeepEqual(CircularNumbers(c.n), c.want) {
			t.Errorf("Failed to check case: %+v\n", c)
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

func TestIsCircularPrime(t *testing.T) {
	cases := []struct {
		n    uint64
		want bool
	}{
		{2, true},
		{79, true},
		{197, true},
		{0, false},
		{10, false},
		{1, false},
	}
	for _, c := range cases {
		if IsCircularPrime(c.n) != c.want {
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

func TestNumberOfPrimes(t *testing.T) {
	cases := []struct {
		below, want uint64
	}{
		{10, 7},
		{100, 27},
	}
	for _, c := range cases {
		if NumberOfPrimes(c.below) != c.want {
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

func TestPow(t *testing.T) {
	cases := []struct {
		x, y *big.Int
		want *big.Int
	}{
		{big.NewInt(1), big.NewInt(1), big.NewInt(1)},
		{big.NewInt(10), big.NewInt(10), big.NewInt(10000000000)},
		{big.NewInt(10), big.NewInt(2), big.NewInt(100)},
	}
	for _, c := range cases {
		if c.want.Cmp(Pow(c.x, c.y)) != 0 {
			t.Errorf("Failed to check case: %+v, %+v\n", c, Pow(c.x, c.y))
		}
	}
}

func TestSumOfSelfPowers(t *testing.T) {
	cases := []struct {
		from, to *big.Int
		want     *big.Int
	}{
		{big.NewInt(1), big.NewInt(10), big.NewInt(10405071317)},
	}
	for _, c := range cases {
		if c.want.Cmp(SumOfSelfPowers(c.from, c.to)) != 0 {
			t.Errorf("Failed to check case: %+v\n", c)
		}
	}
}

func TestNumberOfDigits(t *testing.T) {
	cases := []struct {
		n    *big.Int
		want int
	}{
		{big.NewInt(1234567890), 10},
		{big.NewInt(1), 1},
	}
	for _, c := range cases {
		if c.want != NumberOfDigits(c.n) {
			t.Errorf("Failed to check case: %+v\n", c)
		}
	}
}
