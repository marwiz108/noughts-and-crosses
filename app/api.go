package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// HandleRequests handles URL routes and server start
func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/newgame", createNewGame).Methods("POST")
	router.HandleFunc("/game", getGame).Methods("GET")
	router.HandleFunc("/move", updateBoard).Methods("PUT")
	log.Fatal(http.ListenAndServe(":9000", router))
}
