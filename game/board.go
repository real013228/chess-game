package game

import (
	"github.com/real013228/chess/game/helper"
)

type Board struct {
	pieces      map[helper.Coordinates]Piece
	startingFen string
	moves       []helper.Move
}

func newBoard(fen string) *Board {
	return &Board{startingFen: fen, moves: make([]helper.Move, 0)}
}

func (b *Board) SetPiece(coordinates helper.Coordinates, piece Piece) {
	b.pieces[coordinates] = piece
	piece.SetCoordinates(coordinates)
}

func (b *Board) removePiece(coordinates helper.Coordinates) {
	delete(b.pieces, coordinates)
}

func (b *Board) MovePiece(move helper.Move) {
	piece := b.GetPiece(move.GetFromCoordinate())
	b.removePiece(move.GetFromCoordinate())
	b.SetPiece(move.GetToCoordinate(), piece)
	b.moves = append(b.moves, move)
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

func (b Board) IsSquareAttacked(coordinates helper.Coordinates, color helper.Color) bool {
	pieces := b.GetPiecesByColor(color)
	for _, p := range pieces {
		attackedSquares := p.GetAttackedSquares(b)
		for square := range attackedSquares {
			if square == coordinates {
				return true
			}
		}
	}

	return false
}

func (b Board) GetPiecesByColor(c helper.Color) []Piece {
	res := make([]Piece, 0)
	for _, p := range b.pieces {
		if p.GetColor() == c {
			res = append(res, p)
		}
	}

	return res
}
