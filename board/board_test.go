package board

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFactory_NewBoard(t *testing.T) {
	f := NewFactory(0)
	rows := 10
	columns := 20
	board := f.NewBoard(uuid.New().String(), int64(rows), int64(columns), 10)
	assert.Equal(t, rows*columns, len(board.Cells))
}
