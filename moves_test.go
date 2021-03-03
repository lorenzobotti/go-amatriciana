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

	moves, err := b.slidingMoves(0x40 | 0x04)
	if err != nil || len(moves) != 11 {
		t.Fail()
	}

	moves, err = b.slidingMoves(0x30 | 0x03)
	if err != nil || len(moves) != 5 {
		t.Fail()
	}
}

func TestCrawlingMoves(t *testing.T) {
	b, err := BoardFromFEN(DefaultFEN)
	if err != nil {
		t.Fail()
	}

	b[0x30|0x03] = White | Knight
	moves, err := b.crawlingMoves(0x30 | 0x03)
	if err != nil || len(moves) != 6 {
		t.Fail()
	}
}

func TestIsSquareAttacked(t *testing.T) {
	b, err := BoardFromFEN("8/8/8/1R3r2/8/3q4/8/8 w - - 0 1")
	if err != nil {
		t.Fail()
	}

	if !b.IsSquareAttacked(0x41, 0x23) {
		t.Fail()
	}
}