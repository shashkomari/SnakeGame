package main

import (
	"fmt"

	"github.com/shashkomari/SnakeGame/internal/snakegame"
)

func main() {
	var sg = snakegame.CreateSnakeGame(10, 20)

	go sg.UserControl()
	score := sg.Run()
	fmt.Printf("Game Over!\nScore: %d\n", score)
}
