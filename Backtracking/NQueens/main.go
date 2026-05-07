package main

import "fmt"

func main() {
	nqueensResult := make(chan []int)
	go nqueens(8, nqueensResult)

	i := 0
	for result := range nqueensResult {
		fmt.Printf(" ** Solution %v **\n", i)
		i++
		printNQueensResult(result)
		fmt.Println()
	}
}

func nqueens(n int, solutions chan []int) {
	board := make([]int, n)
	for i := range board {
		board[i] = -1
	}

	backtrack(board, 0, solutions)
	close(solutions)
}

func backtrack(board []int, row int, solutions chan []int) {
	n := len(board)
	if row == n {
		solution := make([]int, n)
		copy(solution, board)
		solutions <- solution
		return
	}
	for col := range n {
		if canPlace(board, row, col) {
			board[row] = col
			backtrack(board, row+1, solutions)
			board[row] = -1
		}
	}
}

func canPlace(board []int, row, col int) bool {
	for i, j := range board[:row] {
		if j == col || j-i == col-row || j+i == col+row {
			return false
		}
	}
	return true
}

func printNQueensResult(s []int) {
	for _, col := range s {
		fmt.Print("|")
		for i := range len(s) {
			if i == col {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
			fmt.Print("|")
		}
		fmt.Println()
	}
}
