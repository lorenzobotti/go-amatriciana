package amatriciana

import (
	"log"
	"testing"
)

func TestXyFromString(t *testing.T) {
	coord := XyFromString("e4")
	if coord != 0x34 {
		log.Fatalf("expected e4, found %#2x", coord)
		t.Fail()
	}
}

func TestSlidingMoves(t *testing.T) {
	b, err := BoardFromFEN(DefaultFEN)
	if err != nil {
		t.Fail()
	}

	b[0x40|0x04] = White | Rook
	b[0x30|0x03] = White | Bishop

	moves := b.slidingMoves(0x40 | 0x04)
	if len(moves) != 11 {
		t.Fail()
	}

	moves = b.slidingMoves(0x30 | 0x03)
	if len(moves) != 5 {
		t.Fail()
	}
}

func TestCrawlingMoves(t *testing.T) {
	b, err := BoardFromFEN(DefaultFEN)
	if err != nil {
		t.Fail()
	}

	b[0x30|0x03] = White | Knight
	moves := b.crawlingMoves(0x30 | 0x03)
	if len(moves) != 6 {
		t.Fail()
	}
}

func TestIsSquareAttacked(t *testing.T) {
	b, err := BoardFromFEN("8/8/8/1R3r2/8/3q4/8/8 w - - 0 1")
	if err != nil {
		t.Fail()
	}

	if !b.isSquareAttacked(XyFromString("b5"), XyFromString("d3")) {
		log.Fatalf("isSquareAttacked(square, by): b5 is not attacked by the queen on d3")
		t.Fail()
	}
}
