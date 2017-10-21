package engine

import (
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"bitbucket.org/Member1221/music-gj-engine/engine/audio"
	"time"
	"github.com/faiface/pixel"
)

type Game struct {
	Name string
	Author string
	Resizable bool
	ContentRoot string
	Window *GameWindow
	SpriteBatch SpriteBatch
	Input InputHandler
	Content *ResourceManager

	fps int
	frames int
	fpstick <-chan time.Time
}

func (ga *Game) ClearColor(r,g,b uint8) {
	ga.Window.win.Clear(color.RGBA{r,g,b, 255})
}

func (ga *Game) GetFPS() int {
	return ga.fps
}

func (g *Game) Run(gl Drawable, rect pixel.Rect) {
	go func() {
		g.fpstick = time.Tick(time.Second)
	}()
	audio.Init(44100)
	g.Content = newresman(g.ContentRoot)
	pixelgl.Run(func () {
		runfunc(g, rect, &gl)
	})
}