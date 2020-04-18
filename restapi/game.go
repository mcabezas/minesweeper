package restapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mcabezas/minesweeper/common"
	"github.com/mcabezas/minesweeper/game"
)

func CreateGameHandler(f *game.Factory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params CreateGameRequest
		if err := common.DecodePost(r, &params); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		game, err := f.CreateGame(params.Rows, params.Columns)
		if err != nil {
			log.Printf("There was an issue with the request %s\n", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		res := &CreateGameResponse{
			GameID:    game.ID,
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	}
}

type CreateGameResponse struct {
	GameID    string `json:"gameID"`
}

type CreateGameRequest struct {
	Rows    int64 `json:"rows"`
	Columns int64 `json:"columns"`
}

