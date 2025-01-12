package main

import (
	"fmt"
	"main/ai"
)

type Game struct {
	board         *Board
	blackScore    int
	whiteScore    int
	currentPlayer string
	mode          int
}

func NewGame(size, mode int) *Game {
	return &Game{
		board:         NewBoard(size),
		currentPlayer: "B",
		mode:          mode,
	}
}

func (g *Game) Play() {
	for {
		g.board.Display()
		fmt.Printf("Ходит %s (Счет: Черные %d, Белые %d)\n", g.currentPlayer, g.blackScore, g.whiteScore)

		var move string
		if g.currentPlayer == "B" || g.mode == 0 {
			fmt.Print("Введите ход (например, B5): ")
			fmt.Scanln(&move)
		} else {
			if g.mode == 1 {
				move = ai.RandomMove(g.board)
			} else if g.mode == 2 {
				move = ai.MonteCarloMove(g.board, "W")
			}
			fmt.Printf("Компьютер делает ход: %s\n", move)
		}

		if err := g.board.ApplyMove(move, g.currentPlayer); err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		captured := g.board.CheckCaptures(g.currentPlayer)
		if g.currentPlayer == "B" {
			g.blackScore += captured
		} else {
			g.whiteScore += captured
		}

		if g.blackScore > 3 {
			fmt.Println("Черные выиграли!")
			break
		} else if g.whiteScore > 3 {
			fmt.Println("Белые выиграли!")
			break
		} else if g.board.IsFull() {
			fmt.Println("Игра окончена. Ничья!")
			break
		}

		if g.currentPlayer == "B" {
			g.currentPlayer = "W"
		} else {
			g.currentPlayer = "B"
		}
	}
}
