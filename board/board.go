package board

import (
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
	// This attribute is used to setup a default percentage of cells having mines in a board
	minesAveragePercentagePerBoardDefault int
}

type Storage interface {
	SaveBoard(board *Board) (string, error)
	GetBoardByGameID(gameID string) (*Board, bool, error)
	DeleteBoard(gameID string) error
	GetCell(gameID string, position Position) (*Cell, bool, error)
	UpdateCell(gameID string, cell *Cell) error
}

func NewFactory() *Factory {
	return &Factory{Storage: newMemory()}
}

func (f *Factory) NewBoard(gameID string, rows, columns int64, minesRate int) *Board {
	cells := make(map[Position]*Cell)
	for row := int64(0); row < rows; row++ {
		for column := int64(0); column < columns; column++ {
			pos := Position{Row: row, Column: column}

			// When the random number is inside the passed rate
			// then the cell will be set up with a mine
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

func (f *Factory) CreateBoard(gameID string, rows, columns int64) (*Board, error) {
	board := f.NewBoard(gameID, rows, columns, f.minesAveragePercentagePerBoardDefault)
	_, err := f.SaveBoard(board)
	return board, err
}