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
	Rows    int64
	Columns int64
	Cells   map[cell.Position]*Cell
}

type Cell struct {
	cell.Position
	cell.Status
	HasMine bool
	RedFlag bool
}

type Factory struct {
	Storage
	// This attribute is used to setup a default percentage of cells having mines in a board
	minesAveragePercentagePerBoardDefault int
}

type Storage interface {
	SaveBoard(board *Board) (string, error)
	GetBoardByGameID(gameID string) (*Board, bool, error)
	GetBoardSizeByGameID(gameID string) (int64, int64, bool, error)
	DeleteBoard(gameID string) error
	GetCell(gameID string, position cell.Position) (*Cell, bool, error)
	UpdateCell(gameID string, cell *Cell) error
}

func NewFactory(minesRatePercentage int) *Factory {
	return &Factory{Storage: newMemory(), minesAveragePercentagePerBoardDefault: minesRatePercentage}
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
		GameID: gameID, BoardID: uuid.NewV4().String(), Cells: cells, Rows: rows, Columns: columns,
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

type ReveleadCell struct {
	Row       int64 `json:"row"`
	Columns   int64 `json:"column"`
	HasMine   bool  `json:"has_mine"`
	NearMines int64 `json:"near_mines"`
}

func (f *Factory) RevealCell(gameID string, rows, columns int64, c *Cell, reviewed map[cell.Position]bool) ([]ReveleadCell, error) {
	revealedCells := []ReveleadCell{}
	if reviewed[c.Position] {
		return revealedCells, nil
	}
	reviewed[c.Position] = true
	c.Status = cell.Revealed
	if err := f.UpdateCell(gameID, c); err != nil {
		return revealedCells, err
	}
	adjacentPositions := GetAdjacentPositions(c.Position, rows, columns)
	nearMines, err := f.countAdjacentMines(gameID, adjacentPositions)
	if err != nil {
		return revealedCells, err
	}

	// Revealing adjacent cells with there is no nearMines
	if nearMines == 0 {
		for _, pos := range adjacentPositions {
			c, found, err := f.GetCell(gameID, pos)
			if err == nil && found {
				cells, err := f.RevealCell(gameID, rows, columns, c, reviewed)
				if err != nil {
					revealedCells = append(revealedCells, cells...)
				}
			}
		}
	}
	revealedCells = append(revealedCells, ReveleadCell{
		Row: c.Row, Columns: c.Column, HasMine: c.HasMine, NearMines: nearMines,
	})
	return revealedCells, nil
}

func (f *Factory) countAdjacentMines(gameID string, adjacentPositions []cell.Position) (int64, error) {
	var nearMines int64
	for _, pos := range adjacentPositions {
		cell, _, err := f.GetCell(gameID, pos)
		if err != nil {
			return 0, err
		}
		if cell.HasMine {
			nearMines++
		}
	}
	return nearMines, nil
}

func GetAdjacentPositions(center cell.Position, rows, cols int64) []cell.Position {
	var row, col int64
	adjacents := []cell.Position{}

	// Adding TOP CELL to adjacent slice
	row = center.Row + 1
	col = center.Column
	if row >= 0 && row < rows && col >= 0 && col < cols {
		adjacents = append(adjacents, cell.Position{Row: row, Column: col})
	}

	// Adding DOWN CELL to adjacent slice
	row = center.Row - 1
	col = center.Column
	if row >= 0 && row < rows && col >= 0 && col < cols {
		adjacents = append(adjacents, cell.Position{Row: row, Column: col})
	}

	// Adding LEFT CELL to adjacent slice
	row = center.Row
	col = center.Column - 1
	if row >= 0 && row < rows && col >= 0 && col < cols {
		adjacents = append(adjacents, cell.Position{Row: row, Column: col})
	}

	// Adding RIGHT CELL to adjacent slice
	row = center.Row
	col = center.Column + 1
	if row >= 0 && row < rows && col >= 0 && col < cols {
		adjacents = append(adjacents, cell.Position{Row: row, Column: col})
	}

	// Adding LEFT-TOP CELL to adjacent slice
	row = center.Row + 1
	col = center.Column - 1
	if row >= 0 && row < rows && col >= 0 && col < cols {
		adjacents = append(adjacents, cell.Position{Row: row, Column: col})
	}

	// Adding RIGHT-TOP CELL to adjacent slice
	row = center.Row + 1
	col = center.Column + 1
	if row >= 0 && row < rows && col >= 0 && col < cols {
		adjacents = append(adjacents, cell.Position{Row: row, Column: col})
	}

	// Adding LEFT-DOWN CELL to adjacent slice
	row = center.Row - 1
	col = center.Column - 1
	if row >= 0 && row < rows && col >= 0 && col < cols {
		adjacents = append(adjacents, cell.Position{Row: row, Column: col})
	}

	// Adding RIGHT-DOWN CELL to adjacent slice
	row = center.Row - 1
	col = center.Column + 1
	if row >= 0 && row < rows && col >= 0 && col < cols {
		adjacents = append(adjacents, cell.Position{Row: row, Column: col})
	}
	return adjacents
}

func (f *Factory) CanFlag(c *Cell) error {
	if c.Status == cell.Unrevealed || !c.RedFlag {
		return nil
	}
	return errors.New("INVALID_ACTION")
}

func (f *Factory) CanRemoveFlag(c *Cell) error {
	if c.Status == cell.Unrevealed || c.RedFlag {
		return nil
	}
	return errors.New("INVALID_ACTION")
}

func (f *Factory) DoFlag(gameID string, cell *Cell) error {
	cell.RedFlag = true
	return f.UpdateCell(gameID, cell)
}

func (f *Factory) RemoveFlag(gameID string, cell *Cell) error {
	cell.RedFlag = false
	return f.UpdateCell(gameID, cell)
}
