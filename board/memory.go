package board

import (
	"errors"
	"sync"

	"github.com/mcabezas/minesweeper/board/cell"
)

type memory struct {
	boards *sync.Map
}

func newMemory() *memory {
	return &memory{
		boards: &sync.Map{},
	}
}

func (m *memory) SaveBoard(board *Board) (string, error) {
	m.boards.Store(board.GameID, board)
	return board.GameID, nil
}

func (m *memory) GetBoardByGameID(gameID string) (*Board, bool, error) {
	board, ok := m.boards.Load(gameID)
	if !ok {
		return &Board{}, false, nil
	}
	return board.(*Board), true, nil
}

func (m *memory) GetBoardSizeByGameID(gameID string) (int64, int64, bool, error) {
	board, ok := m.boards.Load(gameID)
	if !ok {
		return 0, 0, false, nil
	}
	b := board.(*Board)
	return b.Rows, b.Columns, true, nil
}

func (m *memory) DeleteBoard(gameID string) error {
	m.boards.Delete(gameID)
	return nil
}

func (m *memory) GetCell(gameID string, position cell.Position) (*Cell, bool, error) {
	if board, ok := m.boards.Load(gameID); ok {
		board := board.(*Board)
		cell := board.Cells[position]
		return cell, true, nil
	}
	return &Cell{}, false, nil
}

func (m *memory) UpdateCell(gameID string, cell *Cell) error {
	if board, ok := m.boards.Load(gameID); ok {
		board := board.(*Board)
		board.Cells[cell.Position] = cell
		return  nil
	}
	return errors.New("NOT_FOUND")
}

