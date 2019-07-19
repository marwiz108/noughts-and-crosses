package main

import (
	"strconv"
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
	Winner string `json:"winner"`
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
		Winner: "none",
	}

	return g
}

func makeMove(row int, col int) string {
	game.Board[row][col] = game.Player
	switchTurns()

	return "Next player turn"
}

func validateMove(row int, col int) string {
	if (row > 2 || col > 2) {
		return "Board position does not exist."
	}
	if game.Board[row][col] != "-" {
		return "Box already checked. Choose another box."
	}

	// TODO - add in check for winning combos
	// if winning combo, declare winner & endGame
	return "valid"
}

func switchTurns() Game {
	if game.Player == "X" {
		game.Player = "O"
	} else {
		game.Player = "X"
	}

	return game
}

func getGame(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}

func createNewGame(w http.ResponseWriter, r *http.Request) {
	game = newGame()
}

func updateBoard(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	row, err := strconv.Atoi(params["row"][0])
	if err != nil {
		fmt.Println(err)
		return
	}
	col, e := strconv.Atoi(params["col"][0])
	if e != nil {
		fmt.Println(e)
		return
	}
	v := validateMove(row, col)
	if v != "valid" {
		json.NewEncoder(w).Encode(v)
		return
	}
	move := makeMove(row, col)
	json.NewEncoder(w).Encode(move)
	json.NewEncoder(w).Encode(game)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	// for x, a := range game.Board {
	// 	for y, _ := range a {
	// 		fmt.Println(x, y)
	// 	}
	// }

	router.HandleFunc("/newgame", createNewGame).Methods("POST")
	router.HandleFunc("/game", getGame).Methods("GET")
	router.HandleFunc("/move", updateBoard).Methods("PUT")
	log.Fatal(http.ListenAndServe(":9000", router))
}

func main() {
	handleRequests()
}
