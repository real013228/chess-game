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
	return nil
}

func newKing(color helper.Color, coordinates helper.Coordinates) *King {
	return &King{color, coordinates}
}

func (k King) GetName() string {
	return "king"
}
