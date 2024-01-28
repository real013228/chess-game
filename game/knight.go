package game

import "github.com/real013228/chess/game/helper"

type Knight struct {
	color       helper.Color
	coordinates helper.Coordinates
}

func (k Knight) GetColor() helper.Color {
	return k.color
}

func (k Knight) GetCoordinates() helper.Coordinates {
	return k.coordinates
}

func (k *Knight) SetCoordinates(newCoordinates helper.Coordinates) {
	k.coordinates = newCoordinates
}

func (k Knight) GetPieceMoves() map[helper.CoordinatesShift]struct{} {
	return map[helper.CoordinatesShift]struct{}{
		*helper.NewCoordinatesShift(1, 2):   {},
		*helper.NewCoordinatesShift(2, 1):   {},
		*helper.NewCoordinatesShift(2, -1):  {},
		*helper.NewCoordinatesShift(1, -2):  {},
		*helper.NewCoordinatesShift(-2, -1): {},
		*helper.NewCoordinatesShift(-1, -2): {},
		*helper.NewCoordinatesShift(-2, 1):  {},
		*helper.NewCoordinatesShift(-1, 2):  {},
	}
}

func newKnight(color helper.Color, coordinates helper.Coordinates) *Knight {
	return &Knight{color, coordinates}
}

func (k Knight) GetName() string {
	return "knight"
}

func (k Knight) IsSquareAvailableForMove(coordinates helper.Coordinates, board Board) bool {
	return IsSquareAvailable(&k, coordinates, board)
}
func (k Knight) GetAttackedSquares(board Board) map[helper.Coordinates]struct{} {
	return GetAvailableMoveSquares(&k, board)
}
