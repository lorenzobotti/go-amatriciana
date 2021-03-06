package amatriciana

import (
	"testing"
)

func TestDefaultFEN(t *testing.T) {
	board, err := BoardFromFEN(DefaultFEN)
	if err != nil {
		t.Fail()
	}

	if board[0]|White == 0 || board[0]|Pawn == 0 {
		t.Fail()
	}
}

func TestBoardFromFEN(t *testing.T) {
	board, err := BoardFromFEN("8/8/8/1R3r2/8/3q4/8/R7 w - - 0 1")
	if err != nil {
		t.Fail()
	}

	if board[0]&(White|Rook) != White|Rook {
		t.Fail()
	}

	if board[0x41]&(White|Rook) != White|Rook {
		t.Fail()
	}
}

func TestMovesInDirection(t *testing.T) {
	board, err := BoardFromFEN(DefaultFEN)
	if err != nil {
		t.Fail()
	}

	moves := board.MovesInDirection(16, 16)

	if len(moves) != 5 {
		t.Fail()
	}
}
