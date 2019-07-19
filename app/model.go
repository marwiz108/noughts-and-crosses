package app

type board [3][3]string

// Game struct
type Game struct {
	Board board `json:"board"`
	Player string `json:"player"`
	Winner string `json:"winner"`
}
