package game

import (
	"bitbucket.org/Member1221/music-gj-engine/engine"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"math"
	"math/rand"
	"time"
)

var NOTESPR *engine.Sprite

type Note struct {
	GameObject
	noteheight, x float64
	startTime     float64
	col           color.RGBA
	used          bool
	failed        bool
	visibl        int
}

func (g *Note) Init(game *engine.Game, createInfo interface{}) {
	var err error
	if NOTESPR == nil {
		NOTESPR, err = game.Content.LoadSprite("images/NOTE-3.png")
		if err != nil {
			fmt.Println("OMFG PANIC!\n" + err.Error())
		}
	}
	g.startTime = createInfo.(float64)
	g.col = color.RGBA{0, 128 + uint8(rand.Int()%100), uint8(255), 255}
	g.Update(game, 0)
	g.used = false
}

func (g *Note) Update(game *engine.Game, delta time.Duration) {
	sprw := NOTESPR.GetSprite().Picture().Bounds().W()
	b := (float64(g.parent.music.Position()) / g.parent.music.Period(float64(g.parent.Info.BeatsPerMinute), 1))
	g.x = (game.Window.GetBounds().W() + (sprw / 2)) - (b * sprw)
	g.x += (g.startTime * sprw)
	g.x += g.parent.player.pos.X
	if (g.x - (sprw / 2)) < g.parent.player.pos.X+g.parent.player.spr.GetSprite().GetSprite().Picture().Bounds().W() {
		if !g.used {
			g.col = color.RGBA{255, 255, 128 + uint8(rand.Int()%100), 255}
			if game.Input.KeyDownOnce(pixelgl.KeyZ) || game.Input.KeyDown(pixelgl.KeyX) {
				if !g.used {
					calc := int((PointDistance(pixel.V(0, g.parent.player.pos.Y), pixel.V(0, g.parent.bgp.Y+g.noteheight))))
					if calc < 2048 {
						g.parent.IncreasePoints(300)
						g.col = color.RGBA{0, 244, 0, 255}
						g.used = true
					} else {
						g.parent.EndCombo()
					}
				} else {
					g.parent.EndCombo()
				}
			}
		}
	}
	if (g.x - (sprw / 2)) < g.parent.player.pos.X {
		if !g.used && !g.failed {
			g.col = color.RGBA{200, 0, 0, 255}
			g.parent.EndCombo()
			g.failed = true
		}
	}
	if g.used {
		g.visibl++
		if g.visibl < 1 {
			g.visibl = 1
		}
		g.col.A -= 255 - uint8(g.visibl)
	}
	if g.x < -256 {
		if len(g.parent.entities) > 0 {
			g.parent.entities = g.parent.entities[1:]
		}
	}
}

func PointDistance(vec pixel.Vec, vec2 pixel.Vec) float64 {
	return math.Pow(vec.X-vec2.X, 2) + math.Pow(vec.Y-vec2.Y, 2)
}

func (g *Note) Draw(game *engine.Game, delta time.Duration) {
	game.SpriteBatch.Begin()
	game.SpriteBatch.Draw(NOTESPR, pixel.V(g.x, (g.parent.bgp.Y+g.noteheight)), 0, 1+((pulse/5)/10)+float64((g.visibl/255)), g.col, nil)
	game.SpriteBatch.End()
}
