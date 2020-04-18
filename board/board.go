package board

import (
	"log"
	"math/rand"

	uuid "github.com/satori/go.uuid"
)

type Position struct {
	Row    int64
	Column int64
}

type Board struct {
	GameID  string
	BoardID string
	Cells   map[Position]*Cell
}

type Status int

const (
	Invalid = iota
	Revealed
	Unrevealed
)

type Cell struct {
	Position
	Status
	HasMine bool
}

type Factory struct {
	Storage
}

type Storage interface {
	SaveBoard(board *Board) (string, error)
	GetCell(gameID string, position Position) (*Cell, bool, error)
	UpdateCell(gameID string, cell *Cell) error
}

func NewFactory(logger *log.Logger) *Factory {
	return &Factory{Storage: newMemory()}
}

func (f *Factory) NewBoard(gameID string, rows, columns int64, minesRate int) *Board {
	cells := make(map[Position]*Cell)
	for row := int64(0); row < rows; row++ {
		for column := int64(0); column < columns; column++ {
			pos := Position{Row: row, Column: column}
			hasMine := rand.Intn(100) < minesRate
			cells[pos] = &Cell{
				Position: pos, Status: Unrevealed, HasMine: hasMine,
			}
		}
	}
	return &Board{
		GameID: gameID, BoardID: uuid.NewV4().String(), Cells: cells,
	}
}