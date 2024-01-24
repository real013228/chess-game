package game

import "github.com/real013228/chess/game/helper"

type Bishop struct {
	color       helper.Color
	coordinates helper.Coordinates
}

func (b Bishop) GetColor() helper.Color {
	return b.color
}

func (b Bishop) GetCoordinates() helper.Coordinates {
	return b.coordinates
}

func (b *Bishop) SetCoordinates(newCoordinates helper.Coordinates) {
	b.coordinates = newCoordinates
}

func (b Bishop) GetPieceMoves() map[helper.CoordinatesShift]struct{} {
	res := make(map[helper.CoordinatesShift]struct{})
	for i := -7; i <= 7; i++ {
		if i == 0 {
			continue
		}
		res[*helper.NewCoordinatesShift(i, i)] = struct{}{}
		res[*helper.NewCoordinatesShift(i, -i)] = struct{}{}
	}
	return res
}

func newBishop(color helper.Color, coordinates helper.Coordinates) *Bishop {
	return &Bishop{color, coordinates}
}

func (b Bishop) GetName() string {
	return "bishop"
}
