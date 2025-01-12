package main

// стейт
import (
	"errors"
	"fmt"
)

type Board struct {
	size int
	grid [][]string
}

func NewBoard(size int) *Board {
	grid := make([][]string, size)
	for i := range grid {
		grid[i] = make([]string, size)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}
	return &Board{size: size, grid: grid}
}

func (b *Board) Display() {
	fmt.Print("  ")
	for i := 0; i < b.size; i++ {
		fmt.Printf(" %c", 'A'+i)
	}
	fmt.Println()

	for i := 0; i < b.size; i++ {
		fmt.Printf("%2d", i+1)
		for j := 0; j < b.size; j++ {
			fmt.Printf(" %s", b.grid[i][j])
		}
		fmt.Println()
	}
}

func (b *Board) ApplyMove(move string, player string) error {
	if len(move) < 2 {
		return errors.New("некорректный формат хода")
	}

	col := int(move[0] - 'A')
	row := int(move[1] - '1')

	if col < 0 || col >= b.size || row < 0 || row >= b.size {
		return errors.New("ход выходит за пределы доски")
	}
	if b.grid[row][col] != "." {
		return errors.New("клетка уже занята")
	}

	b.grid[row][col] = player
	return nil
}

func (b *Board) CheckCaptures(player string) int {
	opponent := "B"
	if player == "B" {
		opponent = "W"
	}

	visited := make([][]bool, b.size)
	for i := range visited {
		visited[i] = make([]bool, b.size)
	}

	totalCaptured := 0
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.grid[i][j] == opponent && !visited[i][j] {
				captured, stones := b.isCapturedChain(i, j, opponent, visited)
				if captured {
					totalCaptured += len(stones)
					for _, stone := range stones {
						b.grid[stone[0]][stone[1]] = "."
					}
				}
			}
		}
	}

	return totalCaptured
}

func (b *Board) isCapturedChain(row, col int, player string, visited [][]bool) (bool, [][]int) {
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := [][]int{{row, col}}
	visited[row][col] = true
	chain := [][]int{{row, col}}
	captured := true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			newRow := current[0] + dir[0]
			newCol := current[1] + dir[1]

			if newRow < 0 || newRow >= b.size || newCol < 0 || newCol >= b.size {
				continue
			}

			if visited[newRow][newCol] {
				continue
			}

			if b.grid[newRow][newCol] == "." {
				captured = false
			} else if b.grid[newRow][newCol] == player {
				visited[newRow][newCol] = true
				queue = append(queue, []int{newRow, newCol})
				chain = append(chain, []int{newRow, newCol})
			}
		}
	}

	return captured, chain
}

func (b *Board) IsFull() bool {
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.grid[i][j] == "." {
				return false
			}
		}
	}
	return true
}
