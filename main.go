package main

import (
	"fmt"
	"strings"
)

// connect four idea:
// for each turn, a player can either drop a piece in one of the existing columns,
// create a new column on the left or right side of the board,
// or create a new row on the top of the board.
// at start of game, players set the board dimensions and number of pieces in a row to win.
// or can use default values (7 columns, 6 rows, 4 pieces in a row to win).

type Space int

const (
	Empty Space = iota
	Red
	Yellow
)

type Board struct {
	ToWin int
	Grid  [][]Space
}

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for _, x := range board {
		fmt.Printf("%s\n", strings.Join(x, " "))
	}
}
