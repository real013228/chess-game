package main

import (
	"github.com/real013228/chess/game"
)

func main() {
	bf := game.BoardFactory{}
	// fen := "8/8/8/8/3B4/8/8/8 w - - 0 1"
	// fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	fen := "8/q2r4/5b2/8/bR1Q4/8/1B3N2/3N4 w - - 0 1"
	board := bf.FromFEN(fen)
	// renderer := game.Renderer{}
	// renderer.RenderNoPiece(*board)
	g := game.NewGame(*board)
	g.GameLoop()
}
