package engine

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"fmt"
	"time"
)

type GameWindow struct {
	win *pixelgl.Window
	game *Game
	GameLoop Drawable
	title string
}

func runfunc(game *Game, bounds pixel.Rect, gl *Drawable) {
	w, err := NewWindow(game.Name + " by " + game.Author, game.Resizable, game, *gl, bounds)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.game.Window = w
	w.game.Input = InputHandler{w}
	w.game.SpriteBatch = CreateBatch(w)
	if w.GameLoop != nil {
		w.GameLoop.Init(game, nil)
	} else {
		fmt.Println("Warning: no gameloop has been implemented!")
	}
	lastDelta := time.Now()
	for !w.win.Closed() {
		if w.GameLoop != nil {

			dt := time.Since(lastDelta)
			lastDelta = time.Now()
			w.GameLoop.Update(game, dt)

			w.GameLoop.Draw(game, dt)
			w.win.Update()
			game.frames++
			select {
			case <-game.fpstick:
				game.fps = game.frames
				game.frames = 0
			default:

			}
		}
	}
}

func (gw *GameWindow) GetTitle() string {
	return gw.title
}

func (gw *GameWindow) SetTitle(title string) {
	gw.title = title
	gw.win.SetTitle(title)
}

func (gw *GameWindow) GetBounds() pixel.Rect {
	return gw.win.Bounds()
}

func (gw *GameWindow) SetBounds(rect pixel.Rect) {
	gw.win.SetBounds(rect)
}

func (gw *GameWindow) ToggleSmoothness() {
	gw.win.SetSmooth(!gw.win.Smooth())
}

func NewWindow(title string, resizable bool, game *Game, impl Drawable, bounds pixel.Rect) (*GameWindow, error) {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  title,
		Bounds: bounds,
		Resizable: resizable,
		VSync:true,
	})

	if err != nil {
		return nil, err
	}

	return &GameWindow{
		win:win,
		game:game,
		GameLoop:impl,
		title:title,
	}, nil
}
