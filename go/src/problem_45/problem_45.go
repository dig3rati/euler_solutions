package problem_45

import (
	"common"
)

func TriangularAndPentagonalAndHexagonalNumbers(n uint64) {
	t_seq = common.TriangularNumberSeq()
	p_seq = common.PentagonalNumberSeq()
	h_seq = common.HexagonalNumberSeq()

	var intersection uint64
	select {
	case tn := <-t_seq:
	case pn := <-p_seq:
	case hn := <-h_seq:
	}
}
