package main

import (
	"fmt"

	"github.com/real013228/chess/game"
)

func main() {
	bf := game.BoardFactory{}
	// fen := "8/8/8/8/3B4/8/8/8 w - - 0 1"
	// fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	var line string
	fmt.Scanln(&line)
	fen := "2k5/8/8/b7/4n3/8/8/3K4 w - - 0 1"
	board := bf.FromFEN(fen)
	g := game.NewGame(*board)
	g.GameLoop()
}
