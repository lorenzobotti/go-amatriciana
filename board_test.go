package amatriciana

import (
	"fmt"
	"testing"
)

func TestBoardFromFEN(t *testing.T) {
	board, err := BoardFromFEN(DefaultFEN)
	if err != nil {
		t.Fail()
	}

	if board[0]|White == 0 || board[0]|Pawn == 0 {
		t.Fail()
	}
}

func TestMovesInDirection(t *testing.T) {
	board, err := BoardFromFEN(DefaultFEN)
	if err != nil {
		t.Fail()
	}

	moves := board.MovesInDirection(16, 16)

	for _, move := range moves {
		file, rank := move.To.Coords()
		fmt.Printf("file: %d, rank: %d\n", file, rank)
	}
	if len(moves) != 5 {
		fmt.Println("moves: ", len(moves))
		t.Fail()
	}
}
