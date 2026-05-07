package main

import "fmt"

type Pos struct {
	X, Y int
}

func (p Pos) Add(other Pos) Pos {
	return Pos{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

type Board struct {
	size int
	grid [][]int
}

func NewBoard(n int) *Board {
	grid := make([][]int, n)
	for i := range n {
		grid[i] = make([]int, n)
	}

	return &Board{
		size: n,
		grid: grid,
	}
}

func (board *Board) IsValid(pos Pos) bool {
	return pos.X >= 0 && pos.X < board.size && pos.Y >= 0 && pos.Y < board.size
}

func (board *Board) Visit(pos Pos, moveNum int) {
	board.grid[pos.X][pos.Y] = moveNum
}

func (board *Board) Clear(pos Pos) {
	board.grid[pos.X][pos.Y] = 0
}

func (board *Board) IsVisited(pos Pos) bool {
	return board.grid[pos.X][pos.Y] != 0
}

func (board *Board) GetMove(pos Pos) int {
	return board.grid[pos.X][pos.Y]
}

func (board *Board) Print() {
	for i := range board.size {
		for j := range board.size {
			fmt.Printf("%2d ", board.grid[i][j])
		}
		fmt.Println()
	}
}
