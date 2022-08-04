package snakegame

import "math/rand"

const (
	kUp, kDown, kRight, kLeft = 0, 1, 2, 3
	kHeadIndex                = 0
)

func (sg *SnakeGame) initSnake() {
	startCoordinate := point{(len(sg.board) / 2), (len(sg.board[0]) / 2)}

	sg.snake.body = make([]point, 0, 10)
	sg.snake.body = append(sg.snake.body, startCoordinate)

	sg.snake.tail = sg.snake.body[len(sg.snake.body)-1]

	sg.snake.currentDirectional = rand.Intn(4)
}

type snake struct {
	isFoodEaten        bool
	body               []point
	tail               point
	currentDirectional int
}
