package amatriciana

//Xy represents a coordinate. The four most significant bits are the file and
//the four least significant bits are the row
type Xy byte

//Coords separates the file and the rank bits
func (x Xy) Coords() (file byte, rank byte) {
	file = byte(x) & 0x0f
	rank = byte(x) & 0xf0 >> 4
	return
}

func (x Xy) String() string {
	file, rank := x.Coords()
	fileLetter := file + 'a'
	rankLetter := rank + '1'
	return string([]byte{fileLetter, rankLetter})
}

//XyFromString generates an Xy from a string like "e4" or "d6"
func XyFromString(input string) Xy {
	inputBytes := []byte(input)

	rank := byte(inputBytes[0] - 'a')
	file := byte(inputBytes[1]-'1') << 4
	return Xy(rank | file)
}

//Move is a move.
type Move struct {
	From, To Xy
}

func (m Move) String() string {
	return m.From.String() + m.To.String()
}

//Move moves a Move
func (p Position) Move(move Move) {
	p.board[move.To] = p.board[move.From]
	p.board[move.From] = 0x00
}

//for now it's just a proof of concept, no checking for checks, no pawn moves or promotion, no en passant
func (p Position) generateMoves() []Move {
	pawnsDirection := 16
	pawnsRank := 1
	var otherColor Piece = Black
	if p.turn == Black {
		pawnsDirection = -16
		pawnsRank = 6
		otherColor = White
	}
	_ = pawnsDirection

	moves := make([]Move, 0, 32)
	for square, piece := range p.board {
		if (piece & 0xf0) != p.turn {
			continue
		}

		if (piece & 0x0f) == Pawn {
			//check if the square right in front is occupied or not
			if (square+pawnsDirection)&0x88 == 0 && p.board[square+pawnsDirection] == 0 {
				//debugPrintf("pawn can move to %s", Xy(square+pawnsDirection).String())
				moves = append(moves, Move{Xy(square), Xy(square + pawnsDirection)})

				//if the pawn is in the initial position it can move two squares
				if (square&0xf0)>>4 == pawnsRank && p.board[square+(pawnsDirection*2)] == 0 {
					//debugPrintf("pawn can move to %s", Xy(square+(pawnsDirection)*2).String())
					moves = append(moves, Move{Xy(square), Xy(square + (pawnsDirection * 2))})
				}
			}

			//check for captures
			if (square+pawnsDirection+1)&0x88 == 0 &&
				p.board[square+pawnsDirection+1]&0xf0 == otherColor {
				moves = append(moves, Move{Xy(square), Xy(square + pawnsDirection + 1)})
			}

			if (square+pawnsDirection-1)&0x88 == 0 &&
				p.board[square+pawnsDirection-1]&0xf0 == otherColor {
				moves = append(moves, Move{Xy(square), Xy(square + pawnsDirection - 1)})
			}
		}

		if isPieceSliding(piece) {
			moves = append(moves, p.board.slidingMoves(Xy(square))...)
		}

		if isPieceCrawling(piece) {
			moves = append(moves, p.board.crawlingMoves(Xy(square))...)
		}
	}

	return moves
}

func (b Board) movesInDirection(start Xy, dir byte) []Move {
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

func (b Board) slidingMoves(start Xy) []Move {
	startPiece := b[start]
	directions, isSliding := slidingPieceDirections[startPiece&0x0f]
	if !isSliding {
		return nil
	}

	output := make([]Move, 0)

	for _, dir := range directions {
		output = append(output, b.movesInDirection(start, byte(dir))...)
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

//this part is useless, for now
//in the pursuit of fast prototyping, i have decided that
//this will be a king-in-check engine (for now!)
//this is because there are many more optimization decisions i have to make
//and having a full working engine will let me experiment more freely

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
