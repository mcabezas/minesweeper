package restapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mcabezas/minesweeper/board"
	"github.com/mcabezas/minesweeper/board/cell"
)

// Cell godoc
// @Summary Patch a Cell
// @Description Reveal a Cell if it is available for that
// @Tags Cell
// @ID Cell
// @Accept  json
// @Produce  json
// @Success 200 {object} restapi.RevealCellResponse
// @Param game_id path string true "Game ID"
// @Param row path int true "Row"
// @Param column path int true "Columns"
// @Router /games/{game_id}/cells/{row}/{column} [patch]
func RevealCellHandler(f *board.Factory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		gameID, found := vars["gameID"]
		if !found {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		row, err := strconv.ParseInt(vars["row"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		column, err := strconv.ParseInt(vars["column"], 10, 64)
		if !found {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if row < 0 || column < 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		c, found, err := f.GetCell(gameID, cell.Position{Row: row, Column: column})
		if !found {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if err != nil {
			log.Printf("There was an issue with the request %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := f.CanRevealCell(c); err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		rows, columns, _, err := f.GetBoardSizeByGameID(gameID)
		if err := f.CanRevealCell(c); err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		revealedCells, err := f.RevealCell(gameID, rows, columns, c, map[cell.Position]bool{})
		if err != nil {
			log.Printf("There was an issue with the request %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(revealedCells)
	}
}

type RevealCellResponse struct {
	HasMine   bool  `json:"mine"`
	NearMines int64 `json:"near_mines"`
}
