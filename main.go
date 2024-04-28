package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Empty = iota
	X
	O
)

var symbols = map[int]string{
	Empty: " ",
	X:     "❌",
	O:     "⭕",
}

type Board [3][3]int

func (b *Board) Print() {
	fmt.Println("   0   1   2")
	for i := range b {
		fmt.Printf("%d ", i)
		for j := range b[i] {
			fmt.Printf(" %s ", symbols[b[i][j]])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("  -----------")
		}
	}
	fmt.Println()
}

func (b *Board) IsFull() bool {
	for _, row := range b {
		for _, cell := range row {
			if cell == Empty {
				return false
			}
		}
	}
	return true
}

func (b *Board) IsWinner(player int) bool {
	for i := 0; i < 3; i++ {
		if b[i][0] == player && b[i][1] == player && b[i][2] == player {
			return true
		}
		if b[0][i] == player && b[1][i] == player && b[2][i] == player {
			return true
		}
	}
	if b[0][0] == player && b[1][1] == player && b[2][2] == player {
		return true
	}
	if b[0][2] == player && b[1][1] == player && b[2][0] == player {
		return true
	}
	return false
}

func (b *Board) MakeMove(x, y, player int) bool {
	if b[x][y] != Empty {
		return false
	}
	b[x][y] = player
	return true
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var board Board
	currentPlayer := X

	for !board.IsFull() && !board.IsWinner(X) && !board.IsWinner(O) {
		board.Print()
		if currentPlayer == X {
			x, y := NextMove(&board, X)
			board.MakeMove(x, y, X)
			currentPlayer = O
		} else {
			x, y := NextMove(&board, O)
			board.MakeMove(x, y, O)
			currentPlayer = X
		}
	}

	board.Print()
	if board.IsWinner(X) {
		fmt.Println("❌ wins!")
	} else if board.IsWinner(O) {
		fmt.Println("⭕ wins!")
	} else {
		fmt.Println("It's a draw!")
	}
}

