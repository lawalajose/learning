package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	var solve func(row int)
	solve = func(row int) {
		if row == 8 {
			// Print one full solution
			for i := 0; i < 8; i++ {
				z01.PrintRune(rune(board[i] + '1')) // '1' means first column
			}
			z01.PrintRune('\n')
			return
		}
		for col := 0; col < 8; col++ {
			safe := true
			// Check if placing a queen at (row, col) is safe
			for i := 0; i < row; i++ {
				if board[i] == col || // same column
					board[i]-col == row-i || // same diagonal ↘
					col-board[i] == row-i { // same diagonal ↙
					safe = false
					break
				}
			}
			if safe {
				board[row] = col // place queen
				solve(row + 1)   // recursive call for next row
			}
		}
	}
	solve(0) // start from first row
}
