package app

import (
	"fmt"
	"strconv"
	"encoding/json"
	"net/http"
)

func createNewGame(w http.ResponseWriter, r *http.Request) {
	game = newGame()
	json.NewEncoder(w).Encode(game)
}

func getGame(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}

func updateBoard(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	// Convert row param to int
	row, err := strconv.Atoi(params["row"][0])
	if err != nil {
		fmt.Println(err)
		return
	}
	// Convert column param to int
	col, e := strconv.Atoi(params["col"][0])
	if e != nil {
		fmt.Println(e)
		return
	}

	v := validateMove(row, col)
	if v != "valid" {
		json.NewEncoder(w).Encode(v)
		json.NewEncoder(w).Encode(game)
		return
	}
	json.NewEncoder(w).Encode(makeMove(row, col))
	json.NewEncoder(w).Encode(game)
}
