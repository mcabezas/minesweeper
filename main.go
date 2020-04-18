package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "minesweeper", log.LstdFlags|log.Lshortfile)
	serverPort := "5000"
	r := mux.NewRouter()

	logger.Println("http://localhost:" + serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, r))
}
