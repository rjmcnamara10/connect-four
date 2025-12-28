package main

import (
	"errors"
	"fmt"
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

func (s Space) ToString() string {
	switch s {
	case Red:
		return "\033[31m●\033[0m"
	case Yellow:
		return "\033[33m●\033[0m"
	case Empty:
		return "·"
	default:
		return "?"
	}
}

type Board struct {
	ToWin int
	Grid  [][]Space
}

func (b *Board) DropPiece(piece Space, col int) error {
	if piece == Empty {
		return errors.New("attempting to drop empty space")
	}
	if col < 0 || col >= len(b.Grid) {
		return errors.New("column out of bounds")
	}

	for row, space := range b.Grid[col] {
		if space == Empty {
			b.Grid[col][row] = piece
			return nil
		}
	}
	return errors.New("column is full")
}

func (b *Board) Print() {
	rowCount := len(b.Grid[0])
	colCount := len(b.Grid)
	for row := rowCount - 1; row >= 0; row-- {
		for col := range colCount {
			fmt.Print(b.Grid[col][row].ToString())
			if col == colCount-1 {
				fmt.Println()
			} else {
				fmt.Print(" ")
			}
		}
	}
}

func (b *Board) CheckWinner() Space {
	// Check horizontal
	rowCount := len(b.Grid[0])
	colCount := len(b.Grid)
	for row := range rowCount {
		curStreak := 0
		streakSpace := Empty
		for col := range colCount {
			curSpace := b.Grid[col][row]
			switch curSpace {
			case Empty:
				curStreak = 0
				streakSpace = Empty
			case streakSpace:
				curStreak++
				// Check if player has enough spaces in a row to win
				if curStreak >= b.ToWin {
					return streakSpace
				}
			default:
				curStreak = 1
				streakSpace = curSpace
			}
		}
	}

	// For diagonal check, first check if possible to get b.ToWin in a row diagonally
	// (also do this for vertical/horizontal?)

	// If no winner, return Empty
	return Empty
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
	fmt.Println("-------------")
	board.DropPiece(Red, 3)
	board.Print()
	fmt.Println("-------------")
	board.DropPiece(Yellow, 2)
	board.Print()
	fmt.Println("-------------")
	board.DropPiece(Red, 2)
	board.Print()
	fmt.Println("-------------")
	board.DropPiece(Yellow, 0)
	board.DropPiece(Yellow, 0)
	board.DropPiece(Yellow, 0)
	board.DropPiece(Yellow, 0)
	board.DropPiece(Yellow, 0)
	board.DropPiece(Yellow, 0)
	board.DropPiece(Yellow, 0)
	board.DropPiece(Yellow, 0)
	if err := board.DropPiece(Yellow, 0); err != nil {
		fmt.Println("Error:", err)
	}
	board.Print()

	// fmt.Print("\033[31m●\033[0m")
	// fmt.Print("\033[33m●\033[0m")
	// fmt.Println("\033[31m●\033[0m")
	// fmt.Print("\033[33m●\033[0m")
	// fmt.Print("\033[31m●\033[0m")
	// fmt.Println("\033[33m●\033[0m")
	// // fmt.Println("\033[33m⬤\033[0m")
	// fmt.Print("\033[31mO\033[0m")
	// fmt.Print("\033[33mO\033[0m")
	// fmt.Println("\033[31mO\033[0m")
	// fmt.Print("\033[33mO\033[0m")
	// fmt.Print("\033[31mO\033[0m")
	// fmt.Println("\033[33mO\033[0m")
}
