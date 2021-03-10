package amatriciana

var pieceMaterial = map[Piece]float64{
	Pawn:   1.0,
	Knight: 3.0,
	Bishop: 3.4,
	Rook:   5.0,
	Queen:  9.0,
	King:   10000.0,
}

func (p Position) Evaluate() float64 {
	output := 0.0
	for _, piece := range p.board {
		if piece == 0 {
			continue
		}

		actualPiece := piece & 0x0f
		color := piece & 0xf0

		pieceValue := pieceMaterial[actualPiece]
		if color == Black {
			pieceValue *= -1
		}

		output += pieceValue
	}

	return output
}
