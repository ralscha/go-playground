package main

import "strings"

func solveEightQueen() {
	var result [][]string
	solveEightQueenHelper(0, []string{}, &result)
	print(len(result))
	println()
	for _, r := range result {
		for _, s := range r {
			println(s)
		}
		println()
	}
}

func solveEightQueenHelper(row int, board []string, result *[][]string) {
	if row == 8 {
		*result = append(*result, board)
		return
	}
	for col := range 8 {
		if isValid(row, col, board) {
			solveEightQueenHelper(row+1, append(board, getRow(col)), result)
		}
	}
}

func isValid(row, col int, board []string) bool {
	for i := range row {
		if board[i][col] == 'Q' {
			return false
		}
		if col-row+i >= 0 && board[i][col-row+i] == 'Q' {
			return false
		}
		if col+row-i < 8 && board[i][col+row-i] == 'Q' {
			return false
		}
	}
	return true
}

func getRow(col int) string {
	var row strings.Builder
	for i := range 8 {
		if i == col {
			row.WriteString("Q")
		} else {
			row.WriteString(".")
		}
	}
	return row.String()
}
