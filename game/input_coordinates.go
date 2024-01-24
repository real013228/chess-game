package game

import (
	"fmt"

	"github.com/real013228/chess/game/helper"
)

type InputCoordinates struct {
}

func (ic InputCoordinates) input() helper.Coordinates {
	for {
		fmt.Println("Please enter coordinates, ex. a1")
		var line string
		fmt.Scanln(&line)
		if len(line) != 2 {
			fmt.Println("Invalid size")
			continue
		}
		var fileChar byte = line[0]
		var rankChar byte = line[1]
		if fileChar < 65 || fileChar > 72 || rankChar-'0' <= 0 || rankChar-'0' > 8 {
			fmt.Println("Invalid format")
			continue
		}

		return *helper.NewCoordinates(helper.File(fileChar-'A'+1), int(rankChar-'0'))
	}
}

func (ic InputCoordinates) inputCoordinates(color helper.Color, board Board) helper.Coordinates {
	for {
		coordinates := ic.input()
		if board.IsSquareEmpty(coordinates) {
			fmt.Println("Empty square")
			continue
		}
		piece := board.GetPiece(coordinates)
		if piece.GetColor() != color {
			fmt.Println("Wrong color")
			continue
		}

		moveSquares := GetAvailableMoveSquares(piece, board)
		if len(moveSquares) == 0 {
			fmt.Println("Blocked piece")
			continue
		}

		return coordinates
	}
}

func (ic InputCoordinates) inputAvailableSquare(coordinates map[helper.Coordinates]struct{}) helper.Coordinates {
	for {
		fmt.Println("Enter your move for selected piece")
		input := ic.input()
		if _, ok := coordinates[input]; !ok {
			fmt.Println("Not available square")
			continue
		}

		return input
	}
}
