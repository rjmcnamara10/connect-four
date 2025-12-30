package main

import (
	"errors"
	"fmt"
)

type Board struct {
	ToWin int
	Grid  [][]Space
}

func (b *Board) DropPiece(piece Space, col int) error {
	if piece == Empty {
		return errors.New("Attempting to drop empty space")
	}
	if col < 0 || col >= len(b.Grid) {
		return errors.New("Column out of bounds")
	}

	for row, space := range b.Grid[col] {
		if space == Empty {
			b.Grid[col][row] = piece
			return nil
		}
	}
	return errors.New("Column is full")
}

func (b *Board) Print() {
	rowCount := len(b.Grid[0])
	colCount := len(b.Grid)
	fmt.Println("1 2 3 4 5 6 7")
	for row := rowCount - 1; row >= 0; row-- {
		for col := range colCount {
			fmt.Print(b.Grid[col][row].Symbol())
			if col == colCount-1 {
				fmt.Println()
			} else {
				fmt.Print(" ")
			}
		}
	}
	fmt.Println("-------------")
}

func (b *Board) CheckWinner() Space {
	rowCount := len(b.Grid[0])
	colCount := len(b.Grid)

	// Check horizontal
	if colCount >= b.ToWin {
		for row := range rowCount {
			// Reset streak for each row
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
	}

	// Check vertical
	if rowCount >= b.ToWin {
		for col := range colCount {
			// Reset streak for each column
			curStreak := 0
			streakSpace := Empty
			for row := range rowCount {
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
	}

	// For diagonal check, first check if possible to get b.ToWin in a row diagonally
	// min(rowCount, colCount) >= b.toWin

	//   0 1 2 3 4 5 6
	// 0 · · · · · · ·
	// 1 · · · · · · ·
	// 2 · · · · · · ·
	// 3 · · · · · · ·
	// 4 · · · · · · ·
	// 5 · · · · · · ·

	// (0, 2) ... (3, 5) --> (0, rowCount - b.toWin) ... add one to each until
	// (0, 1)

	// (0, 3)
	// (1, 2)
	// (2, 1)
	// (3, 0)
	// --
	// (0, 4)
	// (1, 3)
	// (2, 2)
	// (3, 1)
	// (4, 0)
	// --
	// (0, 5)
	// (1, 4)
	// (2, 3)
	// (3, 2)
	// (4, 1)
	// (5, 0)
	// --
	// (1, 5)
	// (2, 4)
	// (3, 3)
	// (4, 2)
	// (5, 1)
	// (6, 0)
	// --
	// (2, 5)
	// (3, 4)
	// (4, 3)
	// (5, 2)
	// (6, 1)
	// --
	// (3, 5)
	// (4, 4)
	// (5, 3)
	// (6, 2)

	// If no winner, return Empty
	return Empty
}
