package engine

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
)

type InputHandler struct {
	win *GameWindow
}

func (i *InputHandler) KeyDown(b pixelgl.Button) bool {
	return i.win.win.Pressed(b)
}

func (i *InputHandler) KeyDownOnce(b pixelgl.Button) bool {
	return i.win.win.JustPressed(b)
}

func (i *InputHandler) GetMouseScroll() pixel.Vec {
	return i.win.win.MouseScroll()
}

func (i *InputHandler) GetMousePosition() pixel.Vec {
	return i.win.win.MousePosition()
}