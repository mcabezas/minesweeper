package game

import (
	"github.com/google/uuid"
)

type Game struct {
	ID      string
	Rows    int64
	Columns int64
}

func NewGame(rows, columns int64) *Game {
	return &Game{
		ID:      uuid.New().String(),
		Rows:    rows,
		Columns: columns,
	}
}

type Factory struct {
	Storage
}

type Storage interface {
	SaveGame(game *Game) (string, error)
	GetGame(id string) (*Game, bool, error)
}

func NewFactory() *Factory {
	return &Factory{Storage: newMemory()}
}

func (f *Factory) CreateGame(rows, columns int64) (*Game, error) {
	game := NewGame(rows, columns)
	_, err := f.SaveGame(game)
	return game, err
}
