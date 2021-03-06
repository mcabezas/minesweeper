package restapi

import (
	"github.com/gorilla/mux"
	"github.com/mcabezas/minesweeper/board"
	"github.com/mcabezas/minesweeper/game"
)

func SetUpRoutes(gf *game.Factory, bf *board.Factory, r *mux.Router) {
	r.HandleFunc("/games", CreateGameHandler(gf)).Methods("POST")
	r.HandleFunc("/games/{gameID}", GetGameHandler(gf)).Methods("GET")
	r.HandleFunc("/games/{gameID}/cells/{row}/{column}", RevealCellHandler(bf)).Methods("PATCH")
	r.HandleFunc("/games/{gameID}/cells/{row}/{column}/flag", CreateFlagHandler(bf)).Methods("POST")
	r.HandleFunc("/games/{gameID}/cells/{row}/{column}/flag", RemoveFlagHandler(bf)).Methods("DELETE")

}
