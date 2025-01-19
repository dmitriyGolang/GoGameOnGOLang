package ai

import (
	"math/rand"

	"github.com/dmitriyGolang/GoGameOnGOLang/board"
)

func RandomMove(board *board.Board) string {
	var availableMoves []string

	for i := 0; i < board.Size(); i++ {
		for j := 0; j < board.Size(); j++ {
			cellValue, err := board.GetCell(i, j)
			if err != nil {
				return "АА Ошибка стоп 00000"
			}
			if cellValue == "." {
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
