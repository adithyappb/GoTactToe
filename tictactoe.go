package main

import "math/rand"

func NextMove(b *Board, player int) (int, int) {
	var availableMoves [][2]int
	for i := range b {
		for j := range b[i] {
			if b[i][j] == Empty {
				availableMoves = append(availableMoves, [2]int{i, j})
			}
		}
	}

	randIndex := rand.Intn(len(availableMoves))
	return availableMoves[randIndex][0], availableMoves[randIndex][1]
}


