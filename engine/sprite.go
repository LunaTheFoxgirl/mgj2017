package engine

import "github.com/faiface/pixel"
import (
	"github.com/faiface/pixel/text"
	"image/color"
)

type SpriteLoadable interface {
	GetSprite() *pixel.Sprite
}

type Sprite struct {
	spr *pixel.Sprite
}

func (s *Sprite) GetSprite() *pixel.Sprite {
	return s.spr
}

type Spritesheet struct {
	spr      *pixel.Sprite
	rect     [][]pixel.Rect
	selected pixel.Vec
}

func (s *Spritesheet) GetSprite() *Sprite {
	return &Sprite{pixel.NewSprite(s.spr.Picture(), s.rect[int(s.selected.X)][int(s.selected.Y)])}
}

func (s *Spritesheet) Select(x, y int) {
	s.selected = pixel.V(float64(x), float64(y))
}

type SpriteFont struct {
	atlas *text.Atlas
	text *text.Text
}

func (s *SpriteFont) SetText(txt string) {
	s.text = text.New(pixel.V(0,0), s.atlas)
	s.text.Write([]byte(txt))
}

func (s *SpriteFont) ClearText(txt string) {
	s.text = text.New(pixel.V(0,0), s.atlas)
}

func (s *SpriteFont) WriteText(txt string) {
	s.text.Write([]byte(txt))
}

func (s *SpriteFont) SetColor(rgba color.RGBA) {
	s.text.Color = rgba
}

func (s *SpriteFont) GetCharSize() pixel.Vec {
	return s.atlas.Glyph('A').Frame.Size()
}

func (s *SpriteFont) GetSprite() *pixel.Sprite {
	return pixel.NewSprite(s.atlas.Picture(), s.atlas.Picture().Bounds())
}