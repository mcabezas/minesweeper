package game

import (
	"sync"
)

type memory struct {
	games *sync.Map
}

func newMemory() *memory {
	return &memory{
		games: &sync.Map{},
	}
}

func (m *memory) SaveGame(game *Game) (string, error) {
	m.games.Store(game.ID, game)
	return game.ID, nil
}

func (m *memory) GetGame(id string) (*Game, bool, error) {
	if game, ok := m.games.Load(id); ok {
		return game.(*Game), true, nil
	}
	return &Game{}, false, nil
}
