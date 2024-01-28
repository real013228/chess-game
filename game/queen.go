package game

import "github.com/real013228/chess/game/helper"

type Queen struct {
	color       helper.Color
	coordinates helper.Coordinates
}

func (q Queen) GetColor() helper.Color {
	return q.color
}

func (q Queen) GetCoordinates() helper.Coordinates {
	return q.coordinates
}

func (q *Queen) SetCoordinates(newCoordinates helper.Coordinates) {
	q.coordinates = newCoordinates
}

func (q Queen) GetPieceMoves() map[helper.CoordinatesShift]struct{} {
	res := make(map[helper.CoordinatesShift]struct{})
	for i := -7; i <= 7; i++ {
		if i == 0 {
			continue
		}
		res[*helper.NewCoordinatesShift(i, i)] = struct{}{}
		res[*helper.NewCoordinatesShift(i, -i)] = struct{}{}
		res[*helper.NewCoordinatesShift(i, 0)] = struct{}{}
		res[*helper.NewCoordinatesShift(0, i)] = struct{}{}
	}
	return res
}

func newQueen(color helper.Color, coordinates helper.Coordinates) *Queen {
	return &Queen{color, coordinates}
}

func (q Queen) GetName() string {
	return "queen"
}
func (q Queen) IsSquareAvailableForMove(coordinates helper.Coordinates, board Board) bool {
	res := IsSquareAvailable(&q, coordinates, board)
	if res {
		var coordinatesBetween []helper.Coordinates
		if q.GetCoordinates().GetFile() == coordinates.GetFile() {
			coordinatesBetween = GetVerticalCoordinatesBetween(q.GetCoordinates(), coordinates)
		} else if q.GetCoordinates().GetRank() == coordinates.GetRank() {
			coordinatesBetween = GetHorizontalCoordinatesBetween(q.GetCoordinates(), coordinates)
		} else {
			coordinatesBetween = GetDiagonalCoordinatesBetween(q.GetCoordinates(), coordinates)
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

func (q Queen) GetAttackedSquares(board Board) map[helper.Coordinates]struct{} {
	return GetAvailableMoveSquares(&q, board)
}
