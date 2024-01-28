package helper

type GameState int

const (
	ONGOING GameState = iota + 1
	CHECKMATE_TO_WHITE_KING
	CHECKMATE_TO_BLACK_KING
	STALEMATE
)
