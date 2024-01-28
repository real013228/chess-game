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
		var move helper.Move
		if isWhiteToMove {
			move = ic.InputMove(g.board, helper.WHITE, g.renderer)
		} else {
			move = ic.InputMove(g.board, helper.BLACK, g.renderer)
		}
		g.board.MovePiece(move)
		//input
		//move
		//pass move
		isWhiteToMove = !isWhiteToMove
	}
}
