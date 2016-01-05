package problem_47

import (
	"common"
	"fmt"
	"math"
)

func keys(hash map[uint64]int) []uint64 {
	ks := make([]uint64, len(hash))
	i := 0
	for k := range hash {
		ks[i] = k
		i++
	}
	return ks
}

func DistinctPrimeFactors(num_factors int) []uint64 {
	n := uint64(math.Pow10(num_factors))
	var fdpf uint64
	for i := uint64(2); i < n; i++ {
		pfchan := common.PrimeFactors(i)
		pfs := make(map[uint64]int)
		for pf := range pfchan {
			if pfs[pf] != 0 {
				pfs[pf] += 1
			} else {
				pfs[pf] = 1
			}
		}
		ks := keys(pfs)
		fmt.Println(i, ks, fdpf)
		if len(ks) == num_factors {
			if i-fdpf == 1 {
				return []uint64{fdpf, i}
			}
			fdpf = i
		}
	}
	return nil
}
