package main

import (
	"github.com/real013228/chess/game"
)

func main() {
	bf := game.BoardFactory{}
	board := bf.FromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	// renderer := game.Renderer{}
	// renderer.RenderNoPiece(*board)
	g := game.NewGame(*board)
	g.GameLoop()
}
