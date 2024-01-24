package game

import "github.com/real013228/chess/game/helper"

type Pawn struct {
	color       helper.Color
	coordinates helper.Coordinates
}

func (p Pawn) GetColor() helper.Color {
	return p.color
}

func (p Pawn) GetCoordinates() helper.Coordinates {
	return p.coordinates
}

func (p *Pawn) SetCoordinates(newCoordinates helper.Coordinates) {
	p.coordinates = newCoordinates
}

func newPawn(color helper.Color, coordinates helper.Coordinates) *Pawn {
	return &Pawn{color: color, coordinates: coordinates}
}

func (p Pawn) GetName() string {
	return "pawn"
}

func (p Pawn) GetPieceMoves() map[helper.CoordinatesShift]struct{} {
	return nil
}
