package game

import (
	"strings"
	"unicode"

	"github.com/real013228/chess/game/helper"
)

type BoardFactory struct {
}

func (bf BoardFactory) FromFEN(fen string) *Board {
	board := newBoard(fen)
	board.Init()
	parts := strings.Split(fen, " ")
	piecePositions := parts[0]
	fenRows := strings.Split(piecePositions, "/")
	for i := 0; i < len(fenRows); i++ {
		row := fenRows[i]
		rank := 8 - i
		fileIndex := 1
		for j := 0; j < len(row); j++ {
			fenChar := row[j]
			if unicode.IsDigit(rune(fenChar)) {
				fileIndex += int(fenChar - '0')
			} else {
				file := helper.File(fileIndex)
				coordinates := helper.NewCoordinates(file, rank)
				board.SetPiece(*coordinates, NewPiece(fenChar, *coordinates))
				fileIndex++
			}
		}
	}
	return board
}

func (bf BoardFactory) Copy(board Board) *Board {
	clone := bf.FromFEN(board.startingFen)
	for _, m := range board.moves {
		clone.MovePiece(m)
	}

	return clone
}
