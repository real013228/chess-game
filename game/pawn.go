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
	res := make(map[helper.CoordinatesShift]struct{})
	if p.color == helper.WHITE {
		res[*helper.NewCoordinatesShift(0, 1)] = struct{}{}
		if p.coordinates.GetRank() == 2 {
			res[*helper.NewCoordinatesShift(0, 2)] = struct{}{}
		}

		res[*helper.NewCoordinatesShift(1, 1)] = struct{}{}
		res[*helper.NewCoordinatesShift(-1, 1)] = struct{}{}
	} else {
		res[*helper.NewCoordinatesShift(0, -1)] = struct{}{}
		if p.coordinates.GetRank() == 7 {
			res[*helper.NewCoordinatesShift(0, -2)] = struct{}{}
		}

		res[*helper.NewCoordinatesShift(1, -1)] = struct{}{}
		res[*helper.NewCoordinatesShift(-1, -1)] = struct{}{}
	}

	return res
}
func (p Pawn) IsSquareAvailableForMove(coordinates helper.Coordinates, board Board) bool {
	if p.coordinates.GetFile() == coordinates.GetFile() {
		if abs(coordinates.GetRank()-p.coordinates.GetRank()) == 2 {
			sq := GetVerticalCoordinatesBetween(p.coordinates, coordinates)
			return board.IsSquareEmpty(sq[0]) && board.IsSquareEmpty(coordinates)
		}
		return board.IsSquareEmpty(coordinates)
	}

	return !board.IsSquareEmpty(coordinates) && board.GetPiece(coordinates).GetColor() != p.color
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func (p Pawn) GetAttackedSquares(board Board) map[helper.Coordinates]struct{} {
	res := make(map[helper.Coordinates]struct{})
	if p.color == helper.WHITE {
		if p.GetCoordinates().ValidShift(*helper.NewCoordinatesShift(-1, 1)) {
			res[*helper.NewCoordinates(p.coordinates.GetFile()-1, p.coordinates.GetRank()+1)] = struct{}{}
		}
		if p.GetCoordinates().ValidShift(*helper.NewCoordinatesShift(1, 1)) {
			res[*helper.NewCoordinates(p.coordinates.GetFile()+1, p.coordinates.GetRank()+1)] = struct{}{}
		}
	} else {
		if p.GetCoordinates().ValidShift(*helper.NewCoordinatesShift(-1, -1)) {
			res[*helper.NewCoordinates(p.coordinates.GetFile()-1, p.coordinates.GetRank()-1)] = struct{}{}
		}
		if p.GetCoordinates().ValidShift(*helper.NewCoordinatesShift(1, -1)) {
			res[*helper.NewCoordinates(p.coordinates.GetFile()+1, p.coordinates.GetRank()-1)] = struct{}{}
		}
	}
	return res
}
