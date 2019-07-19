package app

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

func makeMove(row int, col int) string {
	game.Board[row][col] = game.Player
	checkWinner()
	if game.Winner == "" {
		return switchTurns()
	}
	return endGame()
}

func validateMove(row int, col int) string {
	if (row > 2 || col > 2) {
		return "Board position does not exist."
	}
	if game.Board[row][col] != "-" {
		return "Box already checked. Choose another box."
	}

	return "valid"
}

func switchTurns() string {
	if game.Player == "X" {
		game.Player = "O"
	} else {
		game.Player = "X"
	}

	return "Next player turn" // technically this is not needed but is good for confirmation in postman
}

func checkWinner() {
	// Horizontal wins
	for x := 0; x < 3; x++ {
		if game.Board[x][0] != "-" && (game.Board[x][0] == game.Player && game.Board[x][1] == game.Player && game.Board[x][2] == game.Player) {
			game.Winner = "Player " + game.Player
			return
		}
	}

	// Vertical wins
	for y := 0; y < 3; y++ {
		if game.Board[0][y] != "-" && (game.Board[0][y] == game.Player && game.Board[1][y] == game.Player && game.Board[2][y] == game.Player) {
			game.Winner = "Player " + game.Player
			return
		}
	}

	// Diagonal wins
	if game.Board[1][1] != "-" && (game.Board[0][0] == game.Player && game.Board[1][1] == game.Player && game.Board[2][2] == game.Player) {
		game.Winner = "Player " + game.Player
		return
	}
	if game.Board[1][1] != "-" && (game.Board[0][2] == game.Player && game.Board[1][1] == game.Player && game.Board[2][0] == game.Player) {
		game.Winner = "Player " + game.Player
		return
	}

	// Check if board is full - NO WINNER
	if fullBoard() == true {
		game.Winner = "NONE"
		return
	}
}

func fullBoard() bool {
	for x, a := range game.Board {
		for y := range a {
			if game.Board[x][y] == "-" {
				return false
			}
			continue
		}
	}
	return true
}

func endGame() string {
	if game.Winner != "NONE" {
		return game.Winner + " wins!"
	}
	return "Game Over with no winner"
}
