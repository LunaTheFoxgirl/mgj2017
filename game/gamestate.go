package game

import (
	"bitbucket.org/Member1221/music-gj-engine/engine"
	"time"
)

type GameStateManager struct {
	state engine.Drawable
	game *engine.Game
}

func (gsm *GameStateManager) Init(game *engine.Game, createInfo interface{}) {
	gsm.game = game
}

func (gsm *GameStateManager) Update(game *engine.Game, delta time.Duration) {
	gsm.state.Update(game, delta)
}

func (gsm *GameStateManager) Draw(game *engine.Game, delta time.Duration) {
	gsm.state.Draw(game, delta)
}

func (gsm *GameStateManager) PushState(id engine.Drawable) *GameStateManager {
	gsm.state = id
	return gsm
}

func (gsm *GameStateManager) SendInit(createInfo interface{}) *GameStateManager {
	gsm.state.Init(gsm.game, createInfo)
	return gsm
}


