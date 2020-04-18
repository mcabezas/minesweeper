package game

import (
	"github.com/google/uuid"
)

type Game struct {
	ID      string
}

func NewGame() *Game {
	return &Game{
		ID: uuid.New().String(),
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
