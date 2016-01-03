package common

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	cases := []struct {
		str, rev string
	}{
		{"", ""},
		{"Test", "tseT"},
		{"123", "321"},
		{"日本語", "語本日"},
	}
	for _, c := range cases {
		if c.rev != ReverseString(c.str) {
			t.Errorf("Failed to check case: %+v\n", c)
		}
	}
}

func TestIsPalidrome(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"Test", false},
		{"1001", true},
		{"liril", true},
		{"Liril", false},
	}
	for _, c := range cases {
		res := IsPalindrome(c.input)
		if res != c.want {
			t.Errorf("Failed to check case: %+v\n", c)
		}
	}
}

func TestSumOfDigits(t *testing.T) {
	cases := []struct {
		n    string
		want uint64
	}{
		{"123", 6},
		{"123456789", 45},
	}
	for _, c := range cases {
		if SumOfDigits(c.n) != c.want {
			t.Errorf("Failed to check case: %+v\n%+v\n", c, SumOfDigits(c.n))
		}
	}
}
