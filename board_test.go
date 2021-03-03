package amatriciana

import (
	"testing"
	"log"
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

	for i, piece := range board {
		if piece != 0 {
			log.Printf("for - piece: %#2x in position %#2x", piece, i)
		}
	}

	if board[0]&(White|Rook) != White|Rook {
		log.Printf("piece in a1: %x", board[0x14])
		t.Fail()
	}

	if board[0x41]&(White|Rook) != White|Rook {
		log.Printf("piece in b5: %x", board[0x14])
		t.Fail()
	}
}

func TestMovesInDirection(t *testing.T) {
	board, err := BoardFromFEN(DefaultFEN)
	if err != nil {
		t.Fail()
	}

	moves := board.MovesInDirection(16, 16)

	/*for _, move := range moves {
		file, rank := move.To.Coords()
		fmt.Printf("file: %d, rank: %d\n", file+1, rank+1)
	}*/
	if len(moves) != 5 {
		t.Fail()
	}
}
