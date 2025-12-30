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
	// 5 · · · · · · ·
	// 4 · · · · · · ·
	// 3 · · · · · · ·
	// 2 · · · · · · ·
	// 1 · · · · · · ·
	// 0 · · · · · · ·

	// g = b.Grid

	// Slant up diagonal
	// g[0][2] -> g[1][3] -> g[2][4] -> g[3][5]
	// g[0][1] -> g[1][2] -> g[2][3] -> g[3][4] -> g[4][5]
	// g[0][0] -> g[1][1] -> g[2][2] -> g[3][3] -> g[4][4] -> g[5][5]
	// g[1][0] -> g[2][1] -> g[3][2] -> g[4][3] -> g[5][4] -> g[6][5]
	// g[2][0] -> g[3][1] -> g[4][2] -> g[5][3] -> g[6][4]
	// g[3][0] -> g[4][1] -> g[5][2] -> g[6][3]

	// Slant down diagonal
	// g[0][3] -> g[1][2] -> g[2][1] -> g[3][0]
	// g[0][4] -> g[1][3] -> g[2][2] -> g[3][1] -> g[4][0]
	// g[0][5] -> g[1][4] -> g[2][3] -> g[3][2] -> g[4][1] -> g[5][0]
	// g[1][5] -> g[2][4] -> g[3][3] -> g[4][2] -> g[5][1] -> g[6][0]
	// g[2][5] -> g[3][4] -> g[4][3] -> g[5][2] -> g[6][1]
	// g[3][5] -> g[4][4] -> g[5][3] -> g[6][2]

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
					fmt.Printf("g[%v][%v] -> ", col+i, row+i)
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
				fmt.Println()
			}
		}

		fmt.Println()

		// Slant down diagonal (\)
		for col := 0; col <= colCount-b.ToWin; col++ {
			for row := b.ToWin - 1; row < rowCount; row++ {
				if col != 0 && row != rowCount-1 {
					continue // only start at left or top edge
				}
				curStreak := 0
				streakSpace := Empty
				for i := 0; col+i < colCount && row-i >= 0; i++ {
					fmt.Printf("g[%v][%v] -> ", col+i, row-i)
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
				fmt.Println()
			}
		}
	}

	// If no winner, return Empty
	return Empty
}
