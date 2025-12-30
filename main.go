package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// connect four idea:
// for each turn, a player can either drop a piece in one of the existing columns,
// create a new column on the left or right side of the board,
// or create a new row on the top of the board.
// at start of game, players set the board dimensions and number of pieces in a row to win.
// or can use default values (7 columns, 6 rows, 4 pieces in a row to win).

func PromptForColumn(player Space) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%v, enter a column number: ", player)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}

		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter an integer.")
			continue
		}

		return number - 1
	}
}

func main() {
	rowCount := 6
	colCount := 7
	toWin := 4

	grid := make([][]Space, colCount)
	for col := range grid {
		grid[col] = make([]Space, rowCount)
	}

	board := Board{
		ToWin: toWin,
		Grid:  grid,
	}
	board.Print()

	players := []Space{Red, Yellow}

	curTurn := 0
	for {
		curPlayer := players[curTurn%len(players)]
		col := PromptForColumn(curPlayer)
		if err := board.DropPiece(curPlayer, col); err != nil {
			fmt.Printf("%v, please try again.\n", err)
			continue
		}

		board.Print()

		if winner := board.CheckWinner(); winner != Empty {
			fmt.Println(winner, "wins!")
			break
		}

		curTurn++
	}
}
