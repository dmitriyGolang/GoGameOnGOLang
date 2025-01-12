package ai

import (
	"ai/main"
	"math/rand"
)

func RandomMove(board *main.Board) string {
	var availableMoves []string

	for i := 0; i < board.Size(); i++ {
		for j := 0; j < board.Size(); j++ {
			if board.GetCell(i, j) == "." {
				move := string('A'+j) + string('1'+i)
				availableMoves = append(availableMoves, move)
			}
		}
	}

	if len(availableMoves) > 0 {
		rand.Seed(int64(len(availableMoves)))
		return availableMoves[rand.Intn(len(availableMoves))]
	}

	return ""
}
