package game

import (
	"fmt"
	"strings"

	"github.com/real013228/chess/game/helper"
)

type Renderer struct {
}

func (r Renderer) Render(board Board, pieceToMove Piece) {
	var availableMoves map[helper.Coordinates]struct{} = make(map[helper.Coordinates]struct{})
	if pieceToMove != nil {
		availableMoves = GetAvailableMoveSquares(pieceToMove, board)
	}
	for rank := 8; rank >= 1; rank-- {
		var b strings.Builder
		for file := helper.A; file <= helper.H; file++ {
			coordinates := *helper.NewCoordinates(file, rank)
			_, ok := availableMoves[coordinates]
			if board.IsSquareEmpty(coordinates) {
				b.WriteString(r.getSpriteEmptySquare(coordinates, ok))
			} else {
				b.WriteString(r.getPieceSprite(board.pieces[coordinates], ok))
			}
		}
		b.WriteString(helper.ANSI_RESET)
		fmt.Println(b.String())
	}
}

func (r Renderer) RenderNoPiece(board Board) {
	r.Render(board, nil)
}

func (r *Renderer) ColorizeSprite(sprite string, pieceColor helper.Color, isSquareDark bool, isHighlighted bool) string {
	var res strings.Builder

	if isHighlighted {
		res.WriteString(helper.ANSI_HIGHLITED_COLOR)
	} else if isSquareDark {
		res.WriteString(helper.ANSI_BLACK_BG_COLOR)
	} else {
		res.WriteString(helper.ANSI_WHITE_BG_COLOR)
	}

	if pieceColor == helper.WHITE {
		res.WriteString(helper.ANSI_WHITE_PIECE_COLOR)
	} else {
		res.WriteString(helper.ANSI_BLACK_PIECE_COLOR)
	}
	res.WriteString(sprite)
	return res.String()
}

func (r Renderer) getSpriteEmptySquare(coordinates helper.Coordinates, highlight bool) string {
	return r.ColorizeSprite("___", helper.BLACK, helper.IsSquareDark(coordinates), highlight)
}

func (r Renderer) getPieceSprite(piece Piece, highlight bool) string {
	return r.ColorizeSprite(r.selectUnicodeSpritePiece(piece), piece.GetColor(), helper.IsSquareDark(piece.GetCoordinates()), highlight)
}

func (r Renderer) selectUnicodeSpritePiece(piece Piece) string {
	switch piece.GetName() {
	case "pawn":
		return " ♙ "
	case "knight":
		return " ♘ "
	case "bishop":
		return " ♗ "
	case "rook":
		return " ♖ "
	case "king":
		return " ♔ "
	case "queen":
		return " ♕ "
	}
	return ""
}
