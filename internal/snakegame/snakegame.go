package snakegame

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/eiannone/keyboard"
)

const (
	kEmptyCell = '.'
	kApple     = 'o'
	kHead      = 'Z'
	kBody      = 'z'
)

func CreateSnakeGame(h int, w int) *SnakeGame {
	rand.Seed(time.Now().UnixNano())

	var sg SnakeGame

	sg.score = 0

	sg.pauseTime = 850
	sg.board = make([][]rune, h)
	for i := range sg.board {
		sg.board[i] = make([]rune, w)
	}

	sg.clean()
	sg.initSnake()
	sg.foodGenerator()
	return &sg
}

type SnakeGame struct {
	board     [][]rune
	snake     snake
	food      point
	pauseTime int
	score     int
}

func (sg *SnakeGame) moving(turn <-chan int) {
	select {
	case sg.snake.currentDirectional = <-turn:
	default:
	}

	for i := len(sg.snake.body) - 1; i > kHeadIndex; i-- {
		sg.snake.body[i] = sg.snake.body[i-1]
	}

	switch sg.snake.currentDirectional {
	case kUp:
		sg.snake.body[kHeadIndex].x--
	case kDown:
		sg.snake.body[kHeadIndex].x++
	case kRight:
		sg.snake.body[kHeadIndex].y++
	case kLeft:
		sg.snake.body[kHeadIndex].y--
	default:
		panic("Variable currentDirectional has an invalid value")
	}

	sg.wallInteraction()
}

func (p *SnakeGame) Run(turn <-chan int, exit <-chan struct{}) int {

	for {
		p.updateBoard()
		p.showBoard()

		select {
		case <-exit:
			return p.score
		default:
		}

		time.Sleep(time.Duration(p.pauseTime) * time.Millisecond)
		p.moving(turn)
		p.foodInteraction()
		p.clean()
	}

}

func (p *SnakeGame) foodGenerator() {
	x := rand.Intn(len(p.board))
	y := rand.Intn(len(p.board[0]))

	if p.board[x][y] == kEmptyCell {
		p.food = point{x, y}
	} else {
		p.foodGenerator()
	}
}

// Add snake and food to board matrix
func (p *SnakeGame) updateBoard() {
	for index := range p.snake.body {
		if index == kHeadIndex {
			p.board[p.snake.body[index].x][p.snake.body[index].y] = kHead
		} else {
			p.board[p.snake.body[index].x][p.snake.body[index].y] = kBody
		}
	}
	p.board[p.food.x][p.food.y] = kApple
}

func (p *SnakeGame) showBoard() {
	for i := range p.board {
		for j := range p.board[i] {
			fmt.Printf("%c", p.board[i][j])
		}
		fmt.Println()
	}
}

func (p *SnakeGame) clean() {
	for i := range p.board {
		for j := range p.board[i] {
			p.board[i][j] = kEmptyCell
		}
	}
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (p *SnakeGame) foodInteraction() {
	if p.snake.isFoodEaten {
		p.score += 1
		p.snake.body = append(p.snake.body, p.snake.tail)
		p.snake.isFoodEaten = false
	}
	if p.board[p.snake.body[kHeadIndex].x][p.snake.body[kHeadIndex].y] == kApple {
		p.snake.isFoodEaten = true
		p.pauseTime += 50
		p.foodGenerator()
	}
	p.snake.tail = p.snake.body[len(p.snake.body)-1]
}

func (p *SnakeGame) wallInteraction() {
	for i := range p.snake.body {
		if p.snake.body[i].x > len(p.board)-1 {
			p.snake.body[i].x -= len(p.board)
		}
		if p.snake.body[i].y > len(p.board[0])-1 {
			p.snake.body[i].y -= len(p.board[0])
		}
		if p.snake.body[i].x < 0 {
			p.snake.body[i].x += len(p.board)
		}
		if p.snake.body[i].y < 0 {
			p.snake.body[i].y += len(p.board[0])
		}
	}
}

type point struct {
	x int
	y int
}

func (p *SnakeGame) UserControl(turn chan<- int, exit chan<- struct{}) {

	keyData, err := keyboard.GetKeys(10)

	if err != nil {
		panic(err)
	}

	for {
		event := <-keyData
		if event.Err != nil {
			panic(event.Err)
		}

		switch event.Key {
		case keyboard.KeyArrowUp:
			turn <- kUp
		case keyboard.KeyArrowRight:
			turn <- kRight
		case keyboard.KeyArrowDown:
			turn <- kDown
		case keyboard.KeyArrowLeft:
			turn <- kLeft
		case keyboard.KeyEsc:
			exit <- struct{}{}
			return
		default:
		}
	}
}
