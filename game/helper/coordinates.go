package helper

type File int

const (
	A File = iota + 1
	B
	C
	D
	E
	F
	G
	H
)

type Coordinates struct {
	file File
	rank int
}

func NewCoordinates(file File, rank int) *Coordinates {
	return &Coordinates{file: file, rank: rank}
}

func IsSquareDark(coordinates Coordinates) bool {
	return (int(coordinates.file)+coordinates.rank)%2 == 0
}

func (c Coordinates) Shift(shift CoordinatesShift) Coordinates {
	return *NewCoordinates(c.file+File(shift.fileShift), c.rank+shift.rankShift)
}

func (c Coordinates) ValidShift(shift CoordinatesShift) bool {
	nFile := int(c.file) + shift.fileShift
	nRank := c.rank + shift.rankShift
	if nFile <= 0 || nFile > 8 || nRank <= 0 || nRank > 8 {
		return false
	}
	return true
}
