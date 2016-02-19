package common

import (
	"testing"
)

func TestTriangularNumberSeq(t *testing.T) {
	cases := []struct {
		want []uint64
	}{
		{[]uint64{1, 3, 6, 10, 15}},
	}
	for _, c := range cases {
		out := TriangularNumberSeq()
		for _, w := range c.want {
			if w != <-out {
				t.Errorf("Failed to check case: %+v\n", c)
			}
		}
	}
}

func TestPentagonalNumberSeq(t *testing.T) {
	cases := []struct {
		want []uint64
	}{
		{[]uint64{1, 5, 12, 22, 35}},
	}
	for _, c := range cases {
		out := PentagonalNumberSeq()
		for _, w := range c.want {
			if w != <-out {
				t.Errorf("Failed to check case: %+v\n", c)
			}
		}
	}
}

func TestHexagonalNumberSeq(t *testing.T) {
	cases := []struct {
		want []uint64
	}{
		{[]uint64{1, 6, 15, 28, 45}},
	}
	for _, c := range cases {
		out := HexagonalNumberSeq()
		for _, w := range c.want {
			if w != <-out {
				t.Errorf("Failed to check case: %+v\n", c)
			}
		}
	}
}
