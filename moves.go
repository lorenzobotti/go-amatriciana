package amatriciana

//Xy represents a coordinate. The four most significant bits are the file and
//the four least significant bits are the row
type Xy byte

//Coords separates the file and the rank bits
func (a Xy) Coords() (file byte, rank byte) {
	file = byte(a) & 0xf0 >> 4
	rank = byte(a) & 0x0f
	return
}

//XyFromString generates an Xy from a string like "e4" or "d6"
func XyFromString(input string) Xy {
	inputBytes := []byte(input)

	rank := (inputBytes[0] - 'a')
	file := (inputBytes[1] - '1') << 4
	return Xy(rank | file)
}

//Move is a move.
type Move struct {
	From, To Xy
}

//MovesInDirection Generates all the moves possible in a certain direction from a starting point
//It stops when it reaches the edge of the board or another piece (if it's a piece of the opposite color,
//it also adds a capture)

//no references yet
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

/*
func (b Board) generateMoves(color Piece) {
	pawnsDirection := 16
	if color == Black {
		pawnsDirection = -16
	}

	moves := make([]Move, 0, 32)
	for _, piece := range b {
		if (piece & 0xf0) != color {
			continue
		}

	}
}
*/

//todo: forse SlidingMoves() e CrawlingMoves() non devono riportare un errore?
//todo: e non devono neanche essere esportate, no?
func (b Board) slidingMoves(start Xy) []Move {
	startPiece := b[start]
	directions, isSliding := slidingPieceDirections[startPiece&0x0f]
	if !isSliding {
		return nil
	}

	output := make([]Move, 0)
	for _, dir := range directions {
		output = append(output, b.MovesInDirection(start, byte(dir))...)
	}

	return output
}

func (b Board) crawlingMoves(start Xy) []Move {
	directions, isCrawling := crawlingPieceDirections[b[start]&0x0f]
	if !isCrawling {
		return nil
	}
	output := make([]Move, 0)
	for _, dir := range directions {
		target := byte(int8(start) + dir)
		if target&0x88 != 0 {
			continue
		}

		if (b[target] & 0xf0) != (b[start] & 0xf0) {
			output = append(output, Move{start, Xy(target)})
		}
	}

	return output
}

func (b Board) isSquareAttacked(square Xy, by Xy) bool {
	attackDelta := (by - square) + 128
	attackerType := pieceToBitflag[Piece(by&0xf)]

	if attackingDeltas[attackDelta]&attackerType != 0 {
		return true
	}

	/*debugPrintf("raw delta: %d", attackDelta)
	debugPrintf("delta: %#2x", attackingDeltas[attackDelta])
	debugPrintf("attacker type: %#2x", attackerType)*/

	return false
}

func isPieceSliding(input Piece) bool {
	actualPiece := input & 0x0f
	return actualPiece == Rook || actualPiece == Bishop || actualPiece == Queen
}

func isPieceCrawling(input Piece) bool {
	actualPiece := input & 0x0f
	return actualPiece == Knight || actualPiece == King
}

//todo: il fatto che questi sono int8 ma il resto del codice usa byte rompe un pò le palle
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

const (
	bitflagPawn   = 0x01
	bitflagKnight = 0x02
	bitflagBishop = 0x04
	bitflagRook   = 0x08
	bitflagQueen  = 0x10
	bitflagKing   = 0x20
)

var pieceToBitflag = map[Piece]int8{
	0:      0,
	Pawn:   bitflagPawn,
	Knight: bitflagKnight,
	Bishop: bitflagBishop,
	Rook:   bitflagRook,
	Queen:  bitflagQueen,
	King:   bitflagKing,
}

var attackingDeltas [257]int8

func init() {
	//fmt.Println("starting attacking deltas")

	for _, dir := range crawlingPieceDirections[Knight] {
		attackingDeltas[int(dir)+128] |= bitflagKnight
	}

	for _, dir := range crawlingPieceDirections[King] {
		attackingDeltas[int(dir)+128] |= bitflagKnight
	}

	for i := -128; i <= 128; i++ {
		if i%16 == 0 {
			attackingDeltas[i+128] |= bitflagRook | bitflagQueen
		}

		if i%17 == 0 || i%15 == 0 {
			attackingDeltas[i+128] = attackingDeltas[i+128] | bitflagBishop | bitflagQueen
		}

		if i >= -8 && i <= 8 {
			attackingDeltas[i+128] |= bitflagRook | bitflagQueen
		}
	}

	/*for i, delta := range attackingDeltas {
		debugPrintf("%d: %#2x", i, delta)
	}*/
}
