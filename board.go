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

	// Print column numbers
	for colNum := 1; colNum < colCount; colNum++ {
		fmt.Printf("%v ", colNum)
	}
	fmt.Println(colCount)

	// Print grid spaces
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

	// Print bottom border
	for colNum := 1; colNum < colCount; colNum++ {
		fmt.Print("--")
	}
	fmt.Println("-")
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

	// Check diagonals
	if colCount >= b.ToWin && rowCount >= b.ToWin {
		// Slant up diagonal (/)
		for col := 0; col <= colCount-b.ToWin; col++ {
			for row := 0; row <= rowCount-b.ToWin; row++ {
				if col != 0 && row != 0 {
					continue // only start at left or bottom edge
				}
				curStreak := 0
				streakSpace := Empty
				for i := 0; col+i < colCount && row+i < rowCount; i++ {
					curSpace := b.Grid[col+i][row+i]
					switch curSpace {
					case Empty:
						curStreak = 0
						streakSpace = Empty
					case streakSpace:
						curStreak++
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

		// Slant down diagonal (\)
		for col := 0; col <= colCount-b.ToWin; col++ {
			for row := b.ToWin - 1; row < rowCount; row++ {
				if col != 0 && row != rowCount-1 {
					continue // only start at left or top edge
				}
				curStreak := 0
				streakSpace := Empty
				for i := 0; col+i < colCount && row-i >= 0; i++ {
					curSpace := b.Grid[col+i][row-i]
					switch curSpace {
					case Empty:
						curStreak = 0
						streakSpace = Empty
					case streakSpace:
						curStreak++
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
	}

	// If no winner, return Empty
	return Empty
}
