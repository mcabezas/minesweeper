package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mcabezas/minesweeper/board"
	_ "github.com/mcabezas/minesweeper/docs"
	"github.com/mcabezas/minesweeper/game"
	"github.com/mcabezas/minesweeper/restapi"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Minesweeper Go Restful API
// @version 1.0
// @description Minesweeper API implementation

// @contact.name Marcelo Cabezas
// @contact.email mcabezas@outlook.com

// @host localhost:5000
// @BasePath /
func main() {
	logger := log.New(os.Stdout, "minesweeper", log.LstdFlags|log.Lshortfile)
	serverPort := "5000"
	r := mux.NewRouter()
	swaggerDoc(r)

	bf := board.NewFactory()
	gf := game.NewFactory(bf)
	restapi.SetUpRoutes(gf, r)

	logger.Println("http://localhost:" + serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, r))
}

func swaggerDoc(r *mux.Router) *mux.Route {
	return r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
}
