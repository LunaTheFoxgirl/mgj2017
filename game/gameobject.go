package game

import (
	"bitbucket.org/Member1221/music-gj-engine/engine"
	"time"
)

type GameObject struct {
	engine.Drawable
	parent *GameWorld
}

func (g *GameObject) Init(game *engine.Game, createInfo interface{}) {

}

func (g *GameObject) Update(game *engine.Game, delta time.Duration) {

}

func (g *GameObject) Draw(game *engine.Game, delta time.Duration) {

}

func (g *GameObject) AssignWorld(gw *GameWorld) {
	g.parent = gw
}