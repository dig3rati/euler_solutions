package problem_47

import (
	"testing"
)

func TestDistinctPrimeFactors(t *testing.T) {
	cases := []struct {
		n    int
		want []uint64
	}{
		{2, []uint64{14, 15}},
		{3, []uint64{644, 645, 646}},
	}
	for _, c := range cases {
		t.Logf("Case: %+v\tGot: %+v\n", c, DistinctPrimeFactors(c.n))
	}
}
