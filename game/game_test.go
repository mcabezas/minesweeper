package game

import (
	"testing"

	"github.com/mcabezas/minesweeper/board"
	"github.com/stretchr/testify/assert"
)

func TestFactory_CreateGame(t *testing.T) {
	bf := board.NewFactory()
	f := NewFactory(bf)
	rows := 10
	columns := 10
	game, err := f.CreateGame(int64(rows), int64(columns))
	assert.Nil(t, err)
	assert.NotNil(t, game)
	board, _, _ := bf.GetBoardByGameID(game.ID)
	assert.Equal(t, rows*columns, len(board.Cells))
}
