package helper

type Move struct {
	from, to Coordinates
}

func NewMove(from, to Coordinates) *Move {
	return &Move{from, to}
}

func (m Move) GetFromCoordinate() Coordinates {
	return m.from
}

func (m Move) GetToCoordinate() Coordinates {
	return m.to
}
