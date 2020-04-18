package restapi

import (
	"github.com/gorilla/mux"
	"github.com/mcabezas/minesweeper/game"
)

func SetUpRoutes(gf *game.Factory,r *mux.Router) {
	r.HandleFunc("/games", CreateGameHandler(gf)).Methods("POST")
	r.HandleFunc("/games/{gameID}", GetGameHandler(gf)).Methods("GET")
}
