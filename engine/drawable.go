package engine

import "time"

type Drawable interface {
	Init(game *Game, createInfo interface{})
	Update(game *Game, delta time.Duration)
	Draw(game *Game, delta time.Duration)
}