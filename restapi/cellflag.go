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
// @Summary Puts a flag into a specific cell
// @Description Put a flag into a specific cell
// @Tags Cell
// @ID Cell
// @Accept  json
// @Produce  json
// @Success 200
// @Success 204
// @Failure 403
// @Failure 500
// @Router /games/{game_id}/cells/{row}/{column}/flag [post]
func CreateFlagHandler(f *board.Factory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parmeters, ok := getParameters(w, r, f)
		if !ok {
			return
		}
		cell, found, err := f.GetCell(parmeters.gameID, cell.Position{Row: parmeters.row, Column: parmeters.column})
		if !found {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := f.CanFlag(cell); err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if err := f.DoFlag(parmeters.gameID, cell); err != nil {
			log.Printf("There was an issue with the request %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// Cell godoc
// @Summary Remove a flag from a specific cell
// @Description Remove a flag from a specific cell
// @Tags Cell
// @ID Cell
// @Accept  json
// @Produce  json
// @Success 200 {object} restapi.FlagResponse
// @Success 204
// @Failure 403
// @Failure 500
// @Router /games/{game_id}/cells/{row}/{column}/flag [delete]
func RemoveFlagHandler(f *board.Factory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parmeters, ok := getParameters(w, r, f)
		if !ok {
			return
		}
		cell, found, err := f.GetCell(parmeters.gameID, cell.Position{Row: parmeters.row, Column: parmeters.column})
		if !found {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := f.CanFlag(cell); err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if err := f.RemoveFlag(parmeters.gameID, cell); err != nil {
			log.Printf("There was an issue with the request %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		res := &FlagResponse{
			RequestID: "requestID",
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)

	}
}

type RemoveFlagResponse struct {
	RequestID string `json:"requestID"`
}

func getParameters(w http.ResponseWriter, r *http.Request, f *board.Factory) (*FlagRequest, bool) {
	vars := mux.Vars(r)
	gameID, found := vars["gameID"]
	if !found {
		w.WriteHeader(http.StatusNoContent)
		return &FlagRequest{}, false
	}
	row, err := strconv.ParseInt(vars["row"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return &FlagRequest{}, false
	}
	column, err := strconv.ParseInt(vars["column"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return &FlagRequest{}, false
	}
	return &FlagRequest{gameID: gameID, row: row, column: column}, true
}

type FlagResponse struct {
	RequestID string `json:"requestID"`
}

type FlagRequest struct {
	gameID string
	row    int64
	column int64
}
