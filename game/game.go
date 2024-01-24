package game

import "github.com/real013228/chess/game/helper"

type Game struct {
	board    Board
	renderer Renderer
}

func NewGame(board Board) *Game {
	return &Game{board: board, renderer: Renderer{}}
}

func (g *Game) GameLoop() {
	var isWhiteToMove bool = true
	for {
		g.renderer.RenderNoPiece(g.board)
		ic := InputCoordinates{}
		var from helper.Coordinates
		if isWhiteToMove {
			from = ic.inputCoordinates(helper.WHITE, g.board)
		} else {
			from = ic.inputCoordinates(helper.BLACK, g.board)
		}

		piece := g.board.GetPiece(from)
		availableMoveSquares := GetAvailableMoveSquares(piece, g.board)

		g.renderer.Render(g.board, piece)
		to := ic.inputAvailableSquare(availableMoveSquares)
		g.board.MovePiece(from, to)
		//input
		//move
		//pass move
		isWhiteToMove = !isWhiteToMove
	}
}
