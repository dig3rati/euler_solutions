package problem_8

func MaxAdjacentDigitOp(digits uint64, src []byte, op func([]byte) uint64) uint64 {
	slen := uint64(len(src))
	var max uint64
	for i := uint64(0); i <= slen-digits; i++ {
		val := op(src[i : i+digits])
		if max < val {
			max = val
		}
	}
	return max
}
