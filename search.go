package amatriciana

func (p Position) BestMove(depth int) Move {
	bestEval := -10000.0
	bestMove := Move{}

	for _, move := range p.generateMoves() {
		cloneBoard := p
		cloneBoard.Move(move)
		moveEval := -cloneBoard.search(depth)
		if moveEval > bestEval {
			bestEval = moveEval
			bestMove = move
		}
	}

	return bestMove
}

func (p Position) search(depth int) float64 {

	if depth == 0 {
		return p.Evaluate()
	}

	moves := p.generateMoves()
	_ = moves

	bestMoveRating := -100000.0
	for _, move := range moves {
		cloneBoard := p
		cloneBoard.Move(move)

		moveEval := -cloneBoard.search(depth - 1)

		staticEval := cloneBoard.Evaluate()
		if staticEval < 300 {
			moveEval = -10000.0
		}

		bestMoveRating = maxFloat(bestMoveRating, moveEval)
	}

	return bestMoveRating
}
