package problem_47

import (
	"reflect"
	"testing"
)

func TestDistinctPrimeFactors(t *testing.T) {
	checks := []struct {
		n      int
		expect []uint64
	}{
		{2, []uint64{14, 15}},
		{3, []uint64{644, 645, 646}},
		{4, []uint64{}},
	}
	for _, c := range checks {
		dfacs := DistinctPrimeFactors(c.n)
		t.Logf("First %+v consecutive integers to have %+v distinct prime factors are %+v\n",
			c.n, c.n, dfacs)
		if !reflect.DeepEqual(c.expect, dfacs) {
			t.Errorf("Expectation failed:\nExpected: %+v\nGot: %+v\n", c.expect, dfacs)
		}
	}
}
