package common

import (
	"math"
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
				t.Errorf("Failed to check case: %+v", c)
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
			t.Errorf("Failed to check case: %+v", c)
		}
	}
}

func TestPrimeSeq(t *testing.T) {
	cases := []struct {
		upto uint64
		want []uint64
	}{
		{7, []uint64{2, 3, 5, 7}},
		{10, []uint64{2, 3, 5, 7}},
	}
	for _, c := range cases {
		out := PrimeSeq(c.upto)
		for _, w := range c.want {
			if w != <-out {
				t.Errorf("Failed to check case: %+v", c)
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
				t.Errorf("Failed to check case: %+v", c)
			}
		}
	}
}
