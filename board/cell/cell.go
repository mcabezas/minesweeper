package cell

type Position struct {
	Row    int64
	Column int64
}

type Status int

const (
	Invalid = iota
	Revealed
	Unrevealed
)

