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

// connect four idea:
// for each turn, a player can either drop a piece in one of the existing columns,
// create a new column on the left or right side of the board,
// or create a new row on the top of the board.
// at start of game, players set the board dimensions and number of pieces in a row to win.
// or can use default values (7 columns, 6 rows, 4 pieces in a row to win).

// Change double digit column numbers to capital letters

func PromptForInteger(userPrompt string, defaultInt, minInt, maxInt int) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(userPrompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}

		input = strings.TrimSpace(input)
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
	rowCount := PromptForInteger(fmt.Sprintf("Enter the number of rows (default: %v): ", defaultRowCount), defaultRowCount, minRowCount, maxRowCount)
	colCount := PromptForInteger(fmt.Sprintf("Enter the number of columns (default: %v): ", defaultColCount), defaultColCount, minColCount, maxColCount)
	toWin := PromptForInteger(fmt.Sprintf("Enter the number of pieces in a row to win (default: %v): ", defaultToWin), defaultToWin, minToWin, maxToWin)

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
		// change to allow arrow keys to add columns/row
		userCol := PromptForInteger(fmt.Sprintf("%v, enter a column number: ", curPlayer), 0, 1, len(board.Grid))
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
