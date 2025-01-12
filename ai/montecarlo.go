package ai

import (
	"gogame/main"
	"math/rand"
)

func MonteCarloMove(board *main.Board, player string, simulations int) string {
	availableMoves := getAvailableMoves(board)
	bestMove := ""
	bestWinRate := -1.0

	for _, move := range availableMoves {
		winCount := 0

		for i := 0; i < simulations; i++ {
			boardCopy := board.Copy()
			boardCopy.ApplyMove(move, player)

			if simulateRandomGame(boardCopy, switchPlayer(player)) == player {
				winCount++
			}
		}

		winRate := float64(winCount) / float64(simulations)
		if winRate > bestWinRate {
			bestWinRate = winRate
			bestMove = move
		}
	}

	return bestMove
}

func simulateRandomGame(board *main.Board, currentPlayer string) string {
	for !board.IsFull() {
		availableMoves := getAvailableMoves(board)
		if len(availableMoves) == 0 {
			break
		}

		move := availableMoves[rand.Intn(len(availableMoves))]
		board.ApplyMove(move, currentPlayer)

		if board.CheckCaptures(currentPlayer) > 3 {
			return currentPlayer
		}

		currentPlayer = switchPlayer(currentPlayer)
	}

	return "" // Ничья
}

func getAvailableMoves(board *main.Board) []string {
	var moves []string
	for i := 0; i < board.Size(); i++ {
		for j := 0; j < board.Size(); j++ {
			if board.GetCell(i, j) == "." {
				moves = append(moves, string('A'+j)+string('1'+i))
			}
		}
	}
	return moves
}

func switchPlayer(player string) string {
	if player == "B" {
		return "W"
	}
	return "B"
}
