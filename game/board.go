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

func IsSquareAvailableForMove(p Piece, coordinates helper.Coordinates, board Board) bool {
	return p.IsSquareAvailableForMove(coordinates, board)
}
