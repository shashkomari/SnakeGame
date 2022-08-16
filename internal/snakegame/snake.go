package snakegame

import "math/rand"

type DirectionalType int

const (
	kUp DirectionalType = iota
	kDown
	kRight
	kLeft
	kLastElement
)

const kHeadIndex = 0

func (sg *SnakeGame) initSnake() {
	startCoordinate := point{(len(sg.board) / 2), (len(sg.board[0]) / 2)}

	sg.snake.body = make([]point, 0, 10)
	sg.snake.body = append(sg.snake.body, startCoordinate)

	sg.snake.currentDirectional = DirectionalType(rand.Int() % int(kLastElement))
}

type snake struct {
	body               []point
	currentDirectional DirectionalType
}
