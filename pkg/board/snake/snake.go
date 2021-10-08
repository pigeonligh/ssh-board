package snake

import "github.com/DyegoCosta/snake-game/snake"

type snakeBoard struct{}

func (b *snakeBoard) Name() string {
	return "Snake"
}

func (b *snakeBoard) RunOrDie(username string) {
	snake.NewGame().Start()
}

func New() *snakeBoard {
	return &snakeBoard{}
}
