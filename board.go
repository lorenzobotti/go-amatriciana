package amatriciana

import (
	"errors"
	"strings"
)

//Piece represents a piece in a byte.
//The four least significant bits represent the piece type
//and the four most significant represent the piece color
type Piece byte

const (
	Pawn   = 0x01
	Knight = 0x02
	Bishop = 0x03
	Rook   = 0x04
	Queen  = 0x05
	King   = 0x06

	White = 0x10
	Black = 0x20
)

var fenPieces map[rune]Piece

func init() {
	fenPieces = map[rune]Piece{
		'p': Pawn | Black,
		'n': Knight | Black,
		'b': Bishop | Black,
		'r': Rook | Black,
		'q': Queen | Black,
		'k': King | Black,
		'P': Pawn | White,
		'N': Knight | White,
		'B': Bishop | White,
		'R': Rook | White,
		'Q': Queen | White,
		'K': King | White,
	}
}

//Board represents a chessboard using the 0x88 method
//https://www.chessprogramming.org/0x88
type Board [128]Piece

//DefaultFEN is the default FEN
const DefaultFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

//BoardFromFEN makes a Board from a FEN
func BoardFromFEN(fen string) (Board, error) {
	elements := strings.Split(fen, " ")
	board := elements[0]

	files := strings.Split(board, "/")
	if len(files) != 8 {
		return Board{}, errors.New("wrong number of files")
	}

	var output Board
	for rankIdx, rank := range files {
		currentFile := 0
		for _, letter := range rank {
			if letter >= '1' && letter <= '8' {
				currentFile += int(letter - '0')
				continue
			}

			piece, isItAPiece := fenPieces[letter]
			if isItAPiece {
				output[rankIdx*16+currentFile] = piece
				currentFile++
			} else {
				return Board{}, errors.New("not a piece something's wrong here")
			}
		}
	}

	return output, nil
}
