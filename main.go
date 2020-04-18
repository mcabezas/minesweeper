package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mcabezas/minesweeper/game"
	"github.com/mcabezas/minesweeper/restapi"
)

func main() {
	logger := log.New(os.Stdout, "minesweeper", log.LstdFlags|log.Lshortfile)
	serverPort := "5000"
	r := mux.NewRouter()

	gf := game.NewFactory()
	restapi.SetUpRoutes(gf, r)

	logger.Println("http://localhost:" + serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, r))
}
