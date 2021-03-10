package amatriciana

import (
	"log"
	"testing"
)

func TestGenerateMoves(t *testing.T) {
	position, err := PositionFromFEN(DefaultFEN)
	if err != nil {
		t.Fail()
	}

	moves := position.generateMoves()
	//not counting pawn moves yet
	if len(moves) != 20 {
		debugPrintf("expected %d moves, found %d\n", 20, len(moves))
		t.Fail()
	}

	position, err = PositionFromFEN("rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 2")
	if err != nil {
		t.Fail()
	}

	moves = position.generateMoves()
	//not counting pawn moves yet
	if len(moves) != 29 {
		debugPrintf("expected %d moves, found %d\n", 29, len(moves))
		t.Fail()
	}

	/*for i, move := range moves {
		pieceMoved := position.board[move.From]
		log.Printf("%d: %s to %s", i+1, pieceString[pieceMoved&0x0f], move.To.String())
	}*/

}

func TestPawnCaptures(t *testing.T) {
	position, err := PositionFromFEN("2k5/8/4pp2/1p1pP3/PPpP4/2P5/8/4K3 b - - 0 1")
	if err != nil {
		t.Fail()
	}

	moves := position.generateMoves()
	//not counting pawn moves yet
	if len(moves) != 8 {
		debugPrintf("expected %d moves, found %d\n", 8, len(moves))
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
