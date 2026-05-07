package main

import (
	"fmt"
)

func main() {
	n := 8
	startPos := Pos{0, 0}
	knightTour(n, startPos)
}

func knightTour(n int, startPos Pos) {
	board := NewBoard(n)
	var moves = []Pos{
		{-2, -1}, {-2, 1},
		{-1, -2}, {-1, 2},
		{1, -2}, {1, 2},
		{2, -1}, {2, 1},
	}

	var backtrack func(pos Pos, moveNum int) bool
	backtrack = func(pos Pos, moveNum int) bool {
		if moveNum >= board.size*board.size {
			return true
		}

		for _, m := range moves {
			next := pos.Add(m)
			if board.IsValid(next) && !board.IsVisited(next) {
				board.Visit(next, moveNum+1)
				if backtrack(next, moveNum+1) {
					return true
				}
				board.Clear(next)
			}
		}
		return false
	}

	board.Visit(startPos, 1)
	if backtrack(startPos, 1) {
		fmt.Println("Tour completed!")
		board.Print()
	} else {
		fmt.Println("No solutions found!")
	}
}
