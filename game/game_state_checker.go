package game

import "github.com/real013228/chess/game/helper"

type GameStateChecker interface {
	Check(board Board, color helper.Color) helper.GameState
}
