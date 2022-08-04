package main

import (
	"fmt"

	"github.com/shashkomari/SnakeGame/internal/snakegame"
)

func main() {
	var sg = snakegame.CreateSnakeGame(10, 20)
	exit := make(chan struct{}, 1)
	turn := make(chan int, 1)

	go sg.UserControl(turn, exit)
	score := sg.Run(turn, exit)
	fmt.Printf("Game Over!\nScore: %d\n", score)
}
