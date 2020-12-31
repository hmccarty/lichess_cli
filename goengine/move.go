package goengine

type Move struct {
	flag Flag
	kingCastle [2]bool
	queenCastle [2]bool
	points int8
	from uint64
	fromBoard Piece
	fromColor Color
	to uint64
	toBoard Piece
	toColor Color
}

func (move *Move) copy() *Move {
	var new *Move = new(Move)
	new.flag = move.flag
	new.kingCastle[WHITE] = move.kingCastle[WHITE]
	new.kingCastle[BLACK] = move.kingCastle[BLACK]
	new.queenCastle[WHITE] = move.queenCastle[WHITE]
	new.queenCastle[BLACK] = move.queenCastle[BLACK]
	new.points = move.points
	new.from = move.from
	new.fromBoard = move.fromBoard
	new.fromColor = move.fromColor
	new.to = move.to
	new.toBoard = move.toBoard
	new.toColor = move.toColor
	return new
}

func (move *Move) ToString() string {
	var fromSqr uint8 = bitScanForward(move.from)
	var toSqr uint8 = bitScanForward(move.to)
	var startRow uint8 = fromSqr / 8
	var startCol uint8 = fromSqr % 8
	var endRow uint8 = toSqr / 8
	var endCol uint8 = toSqr % 8
	return (string((8 - startCol) + ASCII_COL_OFFSET) +
	        string(startRow + ASCII_ROW_OFFSET) + 
			string((8 - endCol) + ASCII_COL_OFFSET) +
			string(endRow + ASCII_ROW_OFFSET))
}

type Flag uint8
const (
	UNKNOWN Flag = iota
	QUIET
	CAPTURE
	KING_SIDE_CASTLE
	QUEEN_SIDE_CASTLE
	EP_CAPTURE
	PROMOTION
)