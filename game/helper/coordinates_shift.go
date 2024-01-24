package helper

type CoordinatesShift struct {
	fileShift int
	rankShift int
}

func (cs CoordinatesShift) GetFileShift() int {
	return cs.fileShift
}

func (cs CoordinatesShift) GetRankShift() int {
	return cs.rankShift
}

func NewCoordinatesShift(fileShift, rankShift int) *CoordinatesShift {
	return &CoordinatesShift{fileShift, rankShift}
}
