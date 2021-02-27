package amatriciana

import "errors"

//Xy represents a coordinate. The four most significant bits are the file and
//the four least significant bits are the row
type Xy byte

//Coords separates the file and the rank bits
func (a Xy) Coords() (file byte, rank byte) {
	file = byte(a) & 0xf0 >> 4
	rank = byte(a) & 0x0f
	return
}

//Move is a move.
type Move struct {
	From, To Xy
}

//MovesInDirection Generates all the moves possible in a certain direction from a starting point
//It stops when it reaches the edge of the board or another piece (if it's a piece of the opposite color,
//it also adds a capture)
func (b Board) MovesInDirection(start Xy, dir byte) []Move {
	startPiece := b[start]
	output := make([]Move, 0)

	for index := byte(start) + dir; (index & 0x88) == 0; index += dir {
		targetSquare := b[index]
		if (targetSquare & 0xf0) != (startPiece & 0xf0) {
			output = append(output, Move{start, Xy(index)})
			if targetSquare&0xf0 != 0 {
				break
			}
		} else {
			break
		}
	}
	return output
}

func (b Board) SlidingMoves(start Xy) ([]Move, error) {
	startPiece := b[start]
	directions, isSliding := slidingPieceDirections[startPiece&0x0f]
	if !isSliding {
		return nil, errors.New("not a sliding piece")
	}

	output := make([]Move, 0)
	for _, dir := range directions {
		output = append(output, b.MovesInDirection(start, byte(dir))...)
	}

	return output, nil
}

var slidingPieceDirections = map[Piece][]int8{
	Rook:   {16, 1, -16, -1},
	Bishop: {17, -17, 15, -15},
	Queen:  {16, 1, -16, -1, 17, -17, 15, -15},
}
var crawlingPieceDirections = map[Piece][]int8{
	Pawn:   {16, 32, 15, 17},
	Knight: {31, 33, 14, 18, -14, -18, -31, -32},
	King:   {16, 1, -16, -1, 17, -17, 15, -15},
}
