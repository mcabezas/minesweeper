package board

import (
	"errors"
	"math/rand"

	"github.com/mcabezas/minesweeper/board/cell"
	uuid "github.com/satori/go.uuid"
)


type Board struct {
	GameID  string
	BoardID string
	Cells   map[cell.Position]*Cell
}

type Cell struct {
	cell.Position
	cell.Status
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
	GetCell(gameID string, position cell.Position) (*Cell, bool, error)
	UpdateCell(gameID string, cell *Cell) error
}

func NewFactory() *Factory {
	return &Factory{Storage: newMemory()}
}

func (f *Factory) NewBoard(gameID string, rows, columns int64, minesRate int) *Board {
	cells := make(map[cell.Position]*Cell)
	for row := int64(0); row < rows; row++ {
		for column := int64(0); column < columns; column++ {
			pos := cell.Position{Row: row, Column: column}

			// When the random number is inside the passed rate
			// then the cell will be set up with a mine
			hasMine := rand.Intn(100) < minesRate
			cells[pos] = &Cell{
				Position: pos, Status: cell.Unrevealed, HasMine: hasMine,
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

func (f *Factory) CanRevealCell(c *Cell) error {
	if c.Status == cell.Unrevealed {
		return nil
	}
	return errors.New("INVALID_ACTION")
}

