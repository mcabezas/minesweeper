package restapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mcabezas/minesweeper/common"
	"github.com/mcabezas/minesweeper/game"
)

// Games godoc
// @Summary Creates a new game
// @Description Creates a new game
// @Tags Games
// @ID Games
// @Accept  json
// @Produce  json
// @Success 201 {object} restapi.CreateGameResponse
// @Param CreateParams body restapi.CreateGameRequest true "Create Game input"
// @Router /games [post]
func CreateGameHandler(f *game.Factory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params CreateGameRequest
		if err := common.DecodePost(r, &params); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := f.CheckGameCreationParameters(params.Rows, params.Columns); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		game, err := f.CreateGame(params.Rows, params.Columns)
		if err != nil {
			log.Printf("There was an issue with the request %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		res := &CreateGameResponse{
			GameID: game.ID,
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	}
}

type CreateGameResponse struct {
	GameID string `json:"gameID"`
}

type CreateGameRequest struct {
	Rows    int64 `json:"rows"`
	Columns int64 `json:"columns"`
}

// Games godoc
// @Summary Game
// @Description Retrieves a Game by ID
// @Tags Games
// @ID Games
// @Accept  json
// @Produce  json
// @Success 200 {object} restapi.GetGameResponse
// @Param game_id path string true "Game ID"
// @Router /games/{game_id} [get]
func GetGameHandler(f *game.Factory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		gameID, found := vars["gameID"]
		if !found {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		game, found, err := f.GetGame(gameID)
		if err != nil {
			log.Printf("There was an issue with the request %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !found {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		res := &GetGameResponse{
			GameID:  game.ID,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

type GetGameResponse struct {
	RequestID string `json:"requestID"`
	GameID    string `json:"gameID"`
}
