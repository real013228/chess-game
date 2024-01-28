package game

import "github.com/real013228/chess/game/helper"

type StalemateChecker struct {
}

func (sc StalemateChecker) Check(board Board, color helper.Color) helper.GameState {
	pieces := board.GetPiecesByColor(color)
	for _, piece := range pieces {
		moves := GetAvailableMoveSquares(piece, board)
		if len(moves) != 0 {
			return helper.ONGOING
		}
	}

	return helper.STALEMATE
}
