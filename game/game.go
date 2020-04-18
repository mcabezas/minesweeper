package game

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mcabezas/minesweeper/board"
)

type Game struct {
	ID      string
	Rows    int64
	Columns int64
	BoardID string
}

func NewGame(uuid, boardID string, rows, columns int64) *Game {
	return &Game{
		ID:      uuid,
		Rows:    rows,
		Columns: columns,
		BoardID: boardID,
	}
}

type Factory struct {
	Storage
	bf *board.Factory
}

type Storage interface {
	SaveGame(game *Game) (string, error)
	GetGame(id string) (*Game, bool, error)
}

func NewFactory(bf *board.Factory) *Factory {
	return &Factory{Storage: newMemory(), bf: bf}
}

func (f *Factory) CreateGame(rows, columns int64) (*Game, error) {
	gameID := uuid.New().String()
	board, err := f.bf.CreateBoard(gameID, rows, columns)
	if err != nil {
		return &Game{}, err
	}
	game := NewGame(gameID, board.BoardID, rows, columns)
	_, err = f.SaveGame(game)
	if err != nil {
		// Manual Rollback assuming there is no RDBM's
		// TODO Can be improved using channels to communicate both packages, I did not had time to implement it
		_ = f.bf.DeleteBoard(gameID)
		return &Game{}, err
	}
	return game, err
}

func (f *Factory) CheckGameCreationParameters(rows, columns int64) error {
	if rows == 0 || columns == 0 {
		return errors.New("INVALID_PARAMETERS")
	}
	return nil
}
