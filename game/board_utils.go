package game

import "github.com/real013228/chess/game/helper"

func GetDiagonalCoordinatesBetween(from, to helper.Coordinates) []helper.Coordinates {
	res := make([]helper.Coordinates, 0)
	var fileShift helper.File
	if to.GetFile() > from.GetFile() {
		fileShift = 1
	} else {
		fileShift = -1
	}

	var rankShift int
	if to.GetRank() > from.GetRank() {
		rankShift = 1
	} else {
		rankShift = -1
	}
	fileIndex := from.GetFile() + fileShift
	rankIndex := from.GetRank() + rankShift
	for fileIndex != to.GetFile() && rankIndex != to.GetRank() {
		res = append(res, *helper.NewCoordinates(fileIndex, rankIndex))
		fileIndex += fileShift
		rankIndex += rankShift
	}
	return res
}

func GetVerticalCoordinatesBetween(from, to helper.Coordinates) []helper.Coordinates {
	res := make([]helper.Coordinates, 0)
	var rankShift int
	if to.GetRank() > from.GetRank() {
		rankShift = 1
	} else {
		rankShift = -1
	}
	rankIndex := from.GetRank() + rankShift
	for rankIndex != to.GetRank() {
		res = append(res, *helper.NewCoordinates(from.GetFile(), rankIndex))
		rankIndex += rankShift
	}
	return res
}

func GetHorizontalCoordinatesBetween(from, to helper.Coordinates) []helper.Coordinates {
	res := make([]helper.Coordinates, 0)
	var fileShift helper.File
	if to.GetFile() > from.GetFile() {
		fileShift = 1
	} else {
		fileShift = -1
	}
	fileIndex := from.GetFile() + fileShift
	for fileIndex != to.GetFile() {
		res = append(res, *helper.NewCoordinates(fileIndex, from.GetRank()))
		fileIndex += fileShift
	}
	return res
}
