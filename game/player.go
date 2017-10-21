package game

import (
	"time"
	"bitbucket.org/Member1221/music-gj-engine/engine"
	"github.com/faiface/pixel"
	"image/color"
	"fmt"
	"github.com/faiface/pixel/pixelgl"
)

type Player struct {
	GameObject
	spr *engine.Spritesheet
	cspr *engine.Sprite
	pos pixel.Vec
	speed float64
	rot float64
}

func (p *Player) Init(game *engine.Game, i interface{}) {
	var err error
	p.spr, err = game.Content.LoadSpriteSheet("images/PLAYER_01.png", [][]pixel.Rect{
		{
			pixel.R(0,64, 128, 128),
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	p.pos = pixel.V(128, game.Window.GetBounds().Center().Y)
	p.spr.Select(0,0)
	p.cspr = p.spr.GetSprite()
}

func (p *Player) Update(game *engine.Game, delta time.Duration) {
	spd := 4.0
	maxspd := 10.0
	if game.Input.KeyDown(pixelgl.KeyLeftShift) {
		spd = 6.0
		maxspd = 15.0
	}
	if game.Input.KeyDown(pixelgl.KeyUp) {
		p.speed += spd
	} else if game.Input.KeyDown(pixelgl.KeyDown) {
		p.speed -= spd
	}
	if p.speed > 1 {
		p.rot = (p.speed/maxspd)/2
		p.speed -= 2
		if p.pos.Y > p.parent.GameBounds.Min.Y {
			p.speed = 0
			p.pos.Y = p.parent.GameBounds.Min.Y + 1
		}
	} else if p.speed < -1 {
		p.rot = -(p.speed/-maxspd)/2
		p.speed += 2
		if p.pos.Y < p.parent.GameBounds.Max.Y {
			p.speed = 0
			p.pos.Y = p.parent.GameBounds.Max.Y - 1
		}
	} else {
		p.rot = 0
		p.speed = 0
	}

	if p.speed > maxspd {
		p.speed = maxspd
	}
	if p.speed < -maxspd {
		p.speed = -maxspd
	}
	p.pos.Y += p.speed
}

func (p *Player) Shoot() {

}

func (p *Player) Draw(game *engine.Game, delta time.Duration) {
	game.SpriteBatch.Begin()
	game.SpriteBatch.Draw(p.cspr, p.pos, p.rot, (0.8+((pulse/5)/5)), color.RGBA{128, 255, 128, 255}, nil)
	game.SpriteBatch.End()
}