package amatriciana

import (
	"errors"
	"strconv"
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

var fenPieces = map[rune]Piece{
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

//Board represents a chessboard using the 0x88 method
//https://www.chessprogramming.org/0x88
type Board [128]Piece

//Position represents everything contained in a FEN string
type Position struct {
	board          Board
	turn           Piece
	whiteCanCastle [2]bool
	blackCanCastle [2]bool
	moveNumber     int
	halfMoveNumber int
	enPassant      Xy
}

//DefaultFEN is the default FEN
const DefaultFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

var (
	ErrFENParamNumber = errors.New("invalid number of parameters in FEN string")
	ErrNotAFENPiece   = errors.New("not a valid FEN piece, number or slash")
	ErrNotAFENColor   = errors.New("the color in the FEN string is not 'w' or 'b'")
)

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
				output[(7-rankIdx)*16+currentFile] = piece
				currentFile++
			} else {
				return Board{}, ErrNotAFENPiece
			}
		}
	}

	return output, nil
}

//PositionFromFEN makes a Position from a FEN
func PositionFromFEN(fen string) (Position, error) {
	output := Position{}

	elements := strings.Split(fen, " ")
	if len(elements) != 6 {
		return Position{}, ErrFENParamNumber
	}

	board, err := BoardFromFEN(fen)
	if err != nil {
		return Position{}, err
	}
	output.board = board

	if elements[1] == "w" {
		output.turn = White
	} else if elements[1] == "b" {
		output.turn = Black
	} else {
		return output, ErrNotAFENColor
	}

	output.whiteCanCastle[0] = stringContains(elements[2], 'K')
	output.whiteCanCastle[1] = stringContains(elements[2], 'Q')
	output.blackCanCastle[0] = stringContains(elements[2], 'k')
	output.blackCanCastle[1] = stringContains(elements[2], 'q')

	if elements[3] != "-" {
		output.enPassant = XyFromString(elements[3])
	}

	output.moveNumber, err = strconv.Atoi(elements[4])
	if err != nil {
		return output, err
	}

	output.halfMoveNumber, err = strconv.Atoi(elements[4])
	if err != nil {
		return output, err
	}

	return output, nil
}
