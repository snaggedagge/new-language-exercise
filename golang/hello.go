package main

import (
	"fmt"
)

const GRID int = 5

type Board struct {
	grid [GRID][GRID]byte
}

func NewBoard() *Board {
	var grid [GRID][GRID]byte
	for x := 0; x < GRID; x++ {
		for y := 0; y < GRID; y++ {
			grid[x][y] = ' '
		}
	}
	return &Board{
		grid: grid,
	}
}

func (b *Board) MakeMove(row, col int, player byte) error {
	if row < 1 || row > GRID || col < 1 || col > GRID {
		return fmt.Errorf("invalid move, position out of range")
	}
	if b.grid[row-1][col-1] != ' ' {
		return fmt.Errorf("invalid move, position already occupied")
	}
	b.grid[row-1][col-1] = player
	return nil
}

func (b *Board) Print() {
	for x := 0; x < GRID; x++ {
		fmt.Print("----")
	}
	fmt.Println()
	for x := 0; x < GRID; x++ {
		for y := 0; y < GRID; y++ {
			fmt.Printf("| %c ", b.grid[x][y])
		}
		fmt.Println("|")
		fmt.Println()
	}
	for x := 0; x < GRID; x++ {
		fmt.Print("----")
	}
	fmt.Println()
}

func main() {
	board := NewBoard()
	var player byte = 'X'
	var winner = false

	for !winner {
		board.Print()
		row, col := getPlayerMove(player)
		err := board.MakeMove(row, col, player)
		if err != nil {
			fmt.Println(err)
			continue
		}
		player = getNextPlayer(player)
	}

	board.Print()
}

func getPlayerMove(player byte) (int, int) {
	var row, col int
	fmt.Printf("Player %c, enter row and column: ", player)
	fmt.Scanln(&row, &col)
	return row, col
}

func getNextPlayer(player byte) byte {
	if player == 'X' {
		return 'O'
	}
	return 'X'
}
