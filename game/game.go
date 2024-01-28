package game

import (
	"fmt"

	"github.com/real013228/chess/game/helper"
)

type Game struct {
	board    Board
	renderer Renderer
	checkers []GameStateChecker
}

func NewGame(board Board) *Game {
	return &Game{board: board, renderer: Renderer{}, checkers: []GameStateChecker{
		StalemateChecker{},
		CheckmateChecker{},
	}}
}

func (g *Game) GameLoop() {
	var colorToMove helper.Color = helper.WHITE
	var state helper.GameState
	state = g.determineGameState(g.board, colorToMove)
	for state == helper.ONGOING {
		g.renderer.RenderNoPiece(g.board)
		ic := InputCoordinates{}
		move := ic.InputMove(g.board, colorToMove, g.renderer)
		g.board.MovePiece(move)
		colorToMove = colorToMove.OppositeColor()
		state = g.determineGameState(g.board, colorToMove)
	}
	g.renderer.RenderNoPiece(g.board)
	fmt.Println("Game ended with state ", state)
}

func (g *Game) determineGameState(board Board, color helper.Color) helper.GameState {
	for _, checker := range g.checkers {
		st := checker.Check(board, color)
		if st != helper.ONGOING {
			return st
		}
	}
	return helper.ONGOING
}
