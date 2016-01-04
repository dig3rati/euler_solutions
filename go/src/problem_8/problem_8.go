package problem_8

func AdjacentDigitOp(digits uint64, src []byte, op func(byte, byte) uint64) uint64 {
	slen := uint64(len(src))
	for i := uint64(0); i <= slen-digits; i += digits {
		var prod uint64
		for j := uint64(0); j < digits-1; j++ {
			r := op(src[j:j+1], src[j+1:j+2])
			if prod < r {
				prod = r
			}
		}
	}
	return prod
}
