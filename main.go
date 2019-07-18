package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type board [3][3]string

// Game struct
type Game struct {
	Board board `json:"board"`
	Player string `json:"player"`
}

// Init games var as a slice of Game struct
var game Game

func newGame() Game {
	g := Game{
		Board: board{
			[3]string{"-", "-", "-"},
			[3]string{"-", "-", "-"},
			[3]string{"-", "-", "-"},
		},
		Player: "X",
	}

	return g
}

func getGame(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}

func createNewGame(w http.ResponseWriter, r *http.Request) {
	//
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	game = newGame()
	// for x, a := range game.Board {
	// 	for y, _ := range a {
	// 		fmt.Println(x, y)
	// 	}
	// }

	router.HandleFunc("/game", getGame).Methods("GET")
	// router.HandleFunc("/move", updateBoard).Methods("POST")
	log.Fatal(http.ListenAndServe(":9000", router))
}

func main() {
	handleRequests()
}
