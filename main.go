package main

import (
	"fmt"
)

func main() {
	fmt.Println("Добро пожаловать в упрощенную версию Атари-го!")
	fmt.Println("")
	fmt.Println("Выберите режим игры:")
	fmt.Println("0 - Игрок против игрока")
	fmt.Println("Против бота:")
	fmt.Println("1 - рандом")
	fmt.Println("2 - монтекарло")

	var mode int
	fmt.Scanln(&mode)

	game := NewGame(9, mode)
	game.Play()
}
