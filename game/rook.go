package game

import "github.com/real013228/chess/game/helper"

type Rook struct {
	color       helper.Color
	coordinates helper.Coordinates
}

func (r Rook) GetColor() helper.Color {
	return r.color
}

func (r Rook) GetCoordinates() helper.Coordinates {
	return r.coordinates
}

func (r *Rook) SetCoordinates(newCoordinates helper.Coordinates) {
	r.coordinates = newCoordinates
}

func (r Rook) GetPieceMoves() map[helper.CoordinatesShift]struct{} {
	res := make(map[helper.CoordinatesShift]struct{})
	for i := -7; i <= 7; i++ {
		if i == 0 {
			continue
		}
		res[*helper.NewCoordinatesShift(i, 0)] = struct{}{}
		res[*helper.NewCoordinatesShift(0, i)] = struct{}{}
	}
	return res
}

func newRook(color helper.Color, coordinates helper.Coordinates) *Rook {
	return &Rook{color, coordinates}
}
func (r Rook) GetName() string {
	return "rook"
}

func (r Rook) IsSquareAvailableForMove(coordinates helper.Coordinates, board Board) bool {
	res := IsSquareAvailable(&r, coordinates, board)
	if res {
		var coordinatesBetween []helper.Coordinates
		if r.GetCoordinates().GetFile() != coordinates.GetFile() {
			coordinatesBetween = GetHorizontalCoordinatesBetween(r.GetCoordinates(), coordinates)
		} else {
			coordinatesBetween = GetVerticalCoordinatesBetween(r.GetCoordinates(), coordinates)
		}
		for _, c := range coordinatesBetween {
			if !board.IsSquareEmpty(c) {
				return false
			}
		}
		return true
	}
	return res
}
