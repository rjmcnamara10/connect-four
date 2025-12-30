package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	defaultRowCount = 6
	defaultColCount = 7
	defaultToWin    = 4
	minRowCount     = 1
	minColCount     = 1
	minToWin        = 2
	maxRowCount     = 20
	maxColCount     = 21
	maxToWin        = 20
)

var reader = bufio.NewReader(os.Stdin)

// connect four idea:
// for each turn, a player can either drop a piece in one of the existing columns,
// create a new column on the left or right side of the board,
// or create a new row on the top of the board.
// at start of game, players set the board dimensions and number of pieces in a row to win.
// or can use default values (7 columns, 6 rows, 4 pieces in a row to win).

// Todo:
// - Move letter <-> number functionality to separate file?
// - Allow arrow keys to add columns/row

func PromptForInput(userPrompt string) string {
	for {
		fmt.Print(userPrompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}

		return strings.TrimSpace(input)
	}
}

func promptForInteger(userPrompt string, defaultInt, minInt, maxInt int) int {
	for {
		input := PromptForInput(fmt.Sprintf("%s (default: %v): ", userPrompt, defaultInt))
		if input == "" {
			return defaultInt
		}

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter an integer.")
			continue
		}

		if number < minInt || number > maxInt {
			fmt.Printf("Enter an integer between %v and %v\n", minInt, maxInt)
			continue
		}

		return number
	}
}

func main() {
	rowCount := promptForInteger("Enter the number of rows", defaultRowCount, minRowCount, maxRowCount)
	colCount := promptForInteger("Enter the number of columns", defaultColCount, minColCount, maxColCount)
	toWin := promptForInteger("Enter the number of pieces in a row to win", defaultToWin, minToWin, maxToWin)

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
		userCol := board.PromptForTurn(curPlayer)
		if err := board.DropPiece(curPlayer, userCol-1); err != nil {
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
