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
	return nil
}

func newQueen(color helper.Color, coordinates helper.Coordinates) *Queen {
	return &Queen{color, coordinates}
}

func (q Queen) GetName() string {
	return "queen"
}
