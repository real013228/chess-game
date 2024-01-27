package game

import (
	"github.com/real013228/chess/game/helper"
)

type Piece interface {
	GetName() string
	GetColor() helper.Color
	GetCoordinates() helper.Coordinates
	SetCoordinates(newCoordinates helper.Coordinates)
	GetPieceMoves() map[helper.CoordinatesShift]struct{}
	IsSquareAvailableForMove(coordinates helper.Coordinates, board Board) bool
}

func NewPiece(piece byte, coordinates helper.Coordinates) Piece {
	switch piece {
	case 'p':
		return newPawn(helper.BLACK, coordinates)
	case 'n':
		return newKnight(helper.BLACK, coordinates)
	case 'b':
		return newBishop(helper.BLACK, coordinates)
	case 'r':
		return newRook(helper.BLACK, coordinates)
	case 'k':
		return newKing(helper.BLACK, coordinates)
	case 'q':
		return newQueen(helper.BLACK, coordinates)
	case 'P':
		return newPawn(helper.WHITE, coordinates)
	case 'N':
		return newKnight(helper.WHITE, coordinates)
	case 'B':
		return newBishop(helper.WHITE, coordinates)
	case 'R':
		return newRook(helper.WHITE, coordinates)
	case 'K':
		return newKing(helper.WHITE, coordinates)
	case 'Q':
		return newQueen(helper.WHITE, coordinates)
	}
	return nil
}

func GetAvailableMoveSquares(p Piece, b Board) map[helper.Coordinates]struct{} {
	res := make(map[helper.Coordinates]struct{})
	for shift := range p.GetPieceMoves() {
		if p.GetCoordinates().ValidShift(shift) {
			nCoordinates := p.GetCoordinates().Shift(shift)
			if IsSquareAvailableForMove(p, nCoordinates, b) {
				res[nCoordinates] = struct{}{}
			}
		}
	}

	return res
}
