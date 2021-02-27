package amatriciana

import (
	"testing"
)

func TestSlidingMoves(t *testing.T) {
	b, err := BoardFromFEN(DefaultFEN)
	if err != nil {
		t.Fail()
	}

	b[0x40|0x04] = White | Rook
	b[0x30|0x03] = White | Bishop

	moves, err := b.SlidingMoves(0x40 | 0x04)
	if err != nil || len(moves) != 11 {
		t.Fail()
	}

	moves, err = b.SlidingMoves(0x30 | 0x03)
	if err != nil || len(moves) != 5 {
		t.Fail()
	}
}
