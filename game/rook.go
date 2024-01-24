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
	return nil
}

func newRook(color helper.Color, coordinates helper.Coordinates) *Rook {
	return &Rook{color, coordinates}
}
func (r Rook) GetName() string {
	return "rook"
}
