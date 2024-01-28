package main

import (
	"fmt"

	"github.com/real013228/chess/game"
)

func main() {
	bf := game.BoardFactory{}
	// fen := "8/8/8/8/3B4/8/8/8 w - - 0 1"
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	var line string
	fmt.Scanln(&line)
	// fen := "3R4/k1r5/8/1P1P4/1PKP4/1PPP4/8/8 w - - 0 1"
	board := bf.FromFEN(fen)
	g := game.NewGame(*board)
	g.GameLoop()
}
