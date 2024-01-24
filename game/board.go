package game

import (
	"github.com/real013228/chess/game/helper"
)

type Board struct {
	pieces map[helper.Coordinates]Piece
}

func (b *Board) SetPiece(coordinates helper.Coordinates, piece Piece) {
	b.pieces[coordinates] = piece
	piece.SetCoordinates(coordinates)
}

func (b *Board) removePiece(coordinates helper.Coordinates) {
	delete(b.pieces, coordinates)
}

func (b *Board) MovePiece(from, to helper.Coordinates) {
	piece := b.GetPiece(from)
	b.removePiece(from)
	b.SetPiece(to, piece)
}

func (b *Board) Init() {
	b.pieces = make(map[helper.Coordinates]Piece)
	// for file := helper.A; file <= helper.H; file++ {
	// 	b.SetPiece(*helper.NewCoordinates(file, 2), NewPiece("pawn", helper.WHITE, *helper.NewCoordinates(file, 2)))
	// 	b.SetPiece(*helper.NewCoordinates(file, 7), NewPiece("pawn", helper.BLACK, *helper.NewCoordinates(file, 7)))
	// }
	// b.SetPiece(*helper.NewCoordinates(helper.A, 1), NewPiece("rook", helper.WHITE, *helper.NewCoordinates(helper.A, 1)))
	// b.SetPiece(*helper.NewCoordinates(helper.A, 8), NewPiece("rook", helper.BLACK, *helper.NewCoordinates(helper.A, 8)))
	// b.SetPiece(*helper.NewCoordinates(helper.H, 1), NewPiece("rook", helper.WHITE, *helper.NewCoordinates(helper.H, 1)))
	// b.SetPiece(*helper.NewCoordinates(helper.H, 8), NewPiece("rook", helper.BLACK, *helper.NewCoordinates(helper.H, 8)))

	// b.SetPiece(*helper.NewCoordinates(helper.B, 1), NewPiece("knight", helper.WHITE, *helper.NewCoordinates(helper.B, 1)))
	// b.SetPiece(*helper.NewCoordinates(helper.B, 8), NewPiece("knight", helper.BLACK, *helper.NewCoordinates(helper.B, 8)))
	// b.SetPiece(*helper.NewCoordinates(helper.G, 1), NewPiece("knight", helper.WHITE, *helper.NewCoordinates(helper.G, 1)))
	// b.SetPiece(*helper.NewCoordinates(helper.G, 8), NewPiece("knight", helper.BLACK, *helper.NewCoordinates(helper.G, 8)))

	// b.SetPiece(*helper.NewCoordinates(helper.C, 1), NewPiece("bishop", helper.WHITE, *helper.NewCoordinates(helper.C, 1)))
	// b.SetPiece(*helper.NewCoordinates(helper.C, 8), NewPiece("bishop", helper.BLACK, *helper.NewCoordinates(helper.C, 8)))
	// b.SetPiece(*helper.NewCoordinates(helper.F, 1), NewPiece("bishop", helper.WHITE, *helper.NewCoordinates(helper.F, 1)))
	// b.SetPiece(*helper.NewCoordinates(helper.F, 8), NewPiece("bishop", helper.BLACK, *helper.NewCoordinates(helper.F, 8)))

	// b.SetPiece(*helper.NewCoordinates(helper.D, 1), NewPiece("queen", helper.WHITE, *helper.NewCoordinates(helper.D, 1)))
	// b.SetPiece(*helper.NewCoordinates(helper.D, 8), NewPiece("queen", helper.BLACK, *helper.NewCoordinates(helper.D, 8)))
	// b.SetPiece(*helper.NewCoordinates(helper.E, 1), NewPiece("king", helper.WHITE, *helper.NewCoordinates(helper.E, 1)))
	// b.SetPiece(*helper.NewCoordinates(helper.E, 8), NewPiece("king", helper.BLACK, *helper.NewCoordinates(helper.E, 8)))
}

func (b Board) IsSquareEmpty(coordinates helper.Coordinates) bool {
	_, ok := b.pieces[coordinates]
	return !ok
}

func (b *Board) GetPiece(coordinates helper.Coordinates) Piece {
	return b.pieces[coordinates]
}

func IsSquareAvailable(p Piece, coordinates helper.Coordinates, board Board) bool {
	return board.IsSquareEmpty(coordinates) || board.GetPiece(coordinates).GetColor() != p.GetColor()
}
