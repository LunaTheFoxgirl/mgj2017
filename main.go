package main

import (
	"bitbucket.org/Member1221/music-gj-engine/engine"
	"time"
	"strconv"
	"github.com/faiface/pixel"
	"bitbucket.org/Member1221/music-gj-engine/game"
	"github.com/faiface/pixel/text"
	"log"
	"fmt"
	"math/rand"
)

func main() {
	gm := engine.Game{
		Name: "Rektangel",
		Author: "Clipsey",
		Resizable: false,
	}
	game := MyRythmGame{}
	gm.Run(game, pixel.R(0,0, 980, 620))
}

type MyRythmGame struct {
}

var f *engine.SpriteFont
var err error

var GameState *game.GameStateManager

func (r MyRythmGame) Init(g *engine.Game, i interface{}) {
	g.Window.ToggleSmoothness()
	f, err = g.Content.CreateSpriteFont("fonts/Bedstead.ttf", 14, text.ASCII...)
	if err != nil {
		log.Fatal(err)
		return
	}
	GameState = new(game.GameStateManager)
	GameState.Init(g, i)
	lvls := game.GetLevels(GameState)
	fmt.Println("Listing levels:")
	for _, lvl := range lvls {
		fmt.Println("	- " + lvl.Info.Name + " by " + lvl.Info.Author + " at index " + strconv.Itoa(lvl.Info.LevelIndex))
	}
	rand.Seed(time.Now().Unix())
	game.LoadLevel(GameState, rand.Int()%6)
}

func (r MyRythmGame) Draw(game *engine.Game, delta time.Duration) {
	game.ClearColor(0,0,0)
	GameState.Draw(game, delta)
	game.SpriteBatch.Begin()
	game.SpriteBatch.DrawString(f, "FPS: " + strconv.Itoa(game.GetFPS()), 0, 1, pixel.V(16,game.Window.GetBounds().H()-f.GetCharSize().Y-16), nil)
	game.SpriteBatch.End()
}

func (r MyRythmGame) Update(g *engine.Game, delta time.Duration) {
	GameState.Update(g, delta)
}