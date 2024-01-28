package game

import "github.com/real013228/chess/game/helper"

type CheckmateChecker struct {
}

func (cc CheckmateChecker) Check(board Board, color helper.Color) helper.GameState {
	var king *King
	pieces := board.GetPiecesByColor(color)
	for _, piece := range pieces {
		if val, ok := piece.(*King); ok {
			king = val
		}
	}
	if !board.IsSquareAttacked(king.coordinates, color.OppositeColor()) {
		return helper.ONGOING
	}
	for _, piece := range pieces {
		availableMoves := GetAvailableMoveSquares(piece, board)
		for coordinates := range availableMoves {
			bf := BoardFactory{}
			clone := bf.Copy(board)
			clone.MovePiece(*helper.NewMove(piece.GetCoordinates(), coordinates))
			var clonedKing *King
			npieces := clone.GetPiecesByColor(color)
			for _, piece := range npieces {
				if val, ok := piece.(*King); ok {
					clonedKing = val
				}
			}
			if !clone.IsSquareAttacked(clonedKing.GetCoordinates(), color.OppositeColor()) {
				return helper.ONGOING
			}
		}
	}

	if color == helper.WHITE {
		return helper.CHECKMATE_TO_WHITE_KING
	}
	return helper.CHECKMATE_TO_BLACK_KING
}
