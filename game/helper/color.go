package helper

type Color string

func (c Color) OppositeColor() Color {
	if c == WHITE {
		return BLACK
	}
	return WHITE
}

const (
	WHITE Color = "white"
	BLACK Color = "black"
)

const (
	ANSI_RESET             = "\u001B[0m"
	ANSI_WHITE_PIECE_COLOR = "\u001B[97m"
	ANSI_BLACK_PIECE_COLOR = "\u001B[38;2;180;90;30m"
	ANSI_WHITE_BG_COLOR    = "\u001B[47m"
	ANSI_BLACK_BG_COLOR    = "\u001B[48;2;40;40;40;38;2;255;255;255m"
	ANSI_HIGHLITED_COLOR   = "\u001B[48;2;128;0;128;38;2;255;255;255m"
)
