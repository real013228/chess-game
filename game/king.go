package game

import "github.com/real013228/chess/game/helper"

type King struct {
	color       helper.Color
	coordinates helper.Coordinates
}

func (k King) GetColor() helper.Color {
	return k.color
}

func (k King) GetCoordinates() helper.Coordinates {
	return k.coordinates
}

func (k *King) SetCoordinates(newCoordinates helper.Coordinates) {
	k.coordinates = newCoordinates
}

func (k King) GetPieceMoves() map[helper.CoordinatesShift]struct{} {
	res := make(map[helper.CoordinatesShift]struct{})
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			res[*helper.NewCoordinatesShift(i, j)] = struct{}{}
		}
	}
	return res
}

func newKing(color helper.Color, coordinates helper.Coordinates) *King {
	return &King{color, coordinates}
}

func (k King) GetName() string {
	return "king"
}
func (k King) IsSquareAvailableForMove(coordinates helper.Coordinates, board Board) bool {
	res := IsSquareAvailable(&k, coordinates, board)
	if res {
		return !board.IsSquareAttacked(coordinates, k.color.OppositeColor())
	}

	return res
}

func (k King) GetAttackedSquares(board Board) map[helper.Coordinates]struct{} {
	res := make(map[helper.Coordinates]struct{})
	for shift := range k.GetPieceMoves() {
		if k.GetCoordinates().ValidShift(shift) {
			nCoordinates := k.GetCoordinates().Shift(shift)
			if IsSquareAvailable(&k, nCoordinates, board) {
				res[nCoordinates] = struct{}{}
			}
		}
	}

	return res
}
