package common

func TriangularNumberSeq() chan uint64 {
	return infiniteSequence(func(n uint64) uint64 {
		return n * (n + 1) / 2
	})
}

func PentagonalNumberSeq() chan uint64 {
	return infiniteSequence(func(n uint64) uint64 {
		return n * (3*n - 1) / 2
	})
}

func HexagonalNumberSeq() chan uint64 {
	return infiniteSequence(func(n uint64) uint64 {
		return n * (2*n - 1)
	})
}

func infiniteSequence(operation func(n uint64) uint64) chan uint64 {
	out := make(chan uint64)
	go func() {
		for i := uint64(1); true; i++ {
			out <- operation(i)
		}
	}()
	return out
}
