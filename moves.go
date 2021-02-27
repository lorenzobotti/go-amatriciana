package amatriciana

import "fmt"

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
func (b Board) MovesInDirection(start Xy, dir byte) []Move {
	startPiece := b[start]
	fmt.Println("startPiece:", startPiece)
	fmt.Println("index and 0x88:", start&0x88)
	output := make([]Move, 0)

	for index := byte(start) + dir; (index & 0x88) == 0; index += dir {
		fmt.Println("index:", index)
		targetSquare := b[index]
		if targetSquare&startPiece&0xf0 == 0 {
			output = append(output, Move{start, Xy(index)})
		} else {
			break
		}
	}
	return output
}
