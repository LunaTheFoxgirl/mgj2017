package engine

import (
	"errors"
	"github.com/faiface/pixel"
	"image/color"
)

type SpriteBatch struct {
	win     *GameWindow
	batches map[*Sprite]*pixel.Batch
	fontbatches map[*SpriteFont]*pixel.Batch
	state   uint8
}

func CreateBatch(window *GameWindow) SpriteBatch {
	return SpriteBatch{
		window,
		make(map[*Sprite]*pixel.Batch),
		make(map[*SpriteFont]*pixel.Batch),
		0,
	}
}

func (sp *SpriteBatch) ClearBatch() {
	sp.batches = make(map[*Sprite]*pixel.Batch)
	sp.fontbatches = make(map[*SpriteFont]*pixel.Batch)
}

func (sp *SpriteBatch) Begin() error {
	if sp.state == 0 {
		for _, b := range sp.batches {
			b.Clear()
		}
		for _, b := range sp.fontbatches {
			b.Clear()
		}
		sp.state++
		return nil
	}
	return errors.New("Current batch has not been ended!")
}

func (sp *SpriteBatch) DrawString(font *SpriteFont, text string, rotation float64, scale float64, vec pixel.Vec, camera *Camera) {
	if _, ok := sp.fontbatches[font]; !ok {
		sp.fontbatches[font] = pixel.NewBatch(&pixel.TrianglesData{}, font.atlas.Picture())
	}
	if camera != nil {
		sp.fontbatches[font].SetMatrix(camera.GetMatrixForPos(vec))
	}
	font.SetText(text)
	font.text.Draw(sp.fontbatches[font], pixel.IM.Rotated(pixel.ZV, rotation).Scaled(pixel.ZV, scale).Moved(vec))
}

func (sp *SpriteBatch) DrawSpriteFont(font *SpriteFont, rotation float64, scale float64, vec pixel.Vec, camera *Camera) {
	if _, ok := sp.fontbatches[font]; !ok {
		sp.fontbatches[font] = pixel.NewBatch(&pixel.TrianglesData{}, font.atlas.Picture())
	}
	if camera != nil {
		sp.fontbatches[font].SetMatrix(camera.GetMatrixForPos(vec))
	}
	font.text.Draw(sp.fontbatches[font], pixel.IM.Rotated(pixel.ZV, rotation).Scaled(pixel.ZV, scale).Moved(vec))
}

func (sp *SpriteBatch) Draw(sprite *Sprite, vec pixel.Vec, rotation float64, scale float64, mask color.RGBA, camera *Camera) {
	if _, ok := sp.batches[sprite]; !ok {
		sp.batches[sprite] = pixel.NewBatch(&pixel.TrianglesData{}, sprite.GetSprite().Picture())
	}
	sp.batches[sprite].SetColorMask(mask)
	if camera != nil {
		sp.batches[sprite].SetMatrix(camera.GetMatrixForPos(vec))
	}
	sprite.GetSprite().Draw(sp.batches[sprite], pixel.IM.Rotated(pixel.ZV, rotation).Scaled(pixel.ZV, scale).Moved(vec))
}

func (sp *SpriteBatch) DrawEx(sprite *Sprite, vec pixel.Vec, rotation float64, scale pixel.Vec, mask color.RGBA, camera *Camera) {
	if _, ok := sp.batches[sprite]; !ok {
		sp.batches[sprite] = pixel.NewBatch(&pixel.TrianglesData{}, sprite.GetSprite().Picture())
	}
	sp.batches[sprite].SetColorMask(mask)
	if camera != nil {
		sp.batches[sprite].SetMatrix(camera.GetMatrixForPos(vec))
	}
	sprite.GetSprite().Draw(sp.batches[sprite], pixel.IM.Rotated(pixel.ZV, rotation).ScaledXY(pixel.ZV, scale).Moved(vec))
}

func (sp *SpriteBatch) DrawExWH(sprite *Sprite, vec pixel.Vec, rotation float64, scale pixel.Vec, scaleTri pixel.Vec, mask color.RGBA, camera *Camera) {
	if _, ok := sp.batches[sprite]; !ok {
		sp.batches[sprite] = pixel.NewBatch(&pixel.TrianglesData{
			{
				Position: pixel.Vec{0,0},
			},
			{
				Position: pixel.Vec{0,(1*scaleTri.Y)},
			},
			{
				Position: pixel.Vec{(1*scaleTri.X),(1*scaleTri.Y)},
			},
			{
				Position: pixel.Vec{(1*scaleTri.X),(1*scaleTri.Y)},
			},
			{
				Position: pixel.Vec{0,(1*scaleTri.Y)},
			},
			{
				Position: pixel.Vec{(1*scaleTri.X),0},
			},
		}, sprite.GetSprite().Picture())
	}
	sp.batches[sprite].SetColorMask(mask)
	if camera != nil {
		sp.batches[sprite].SetMatrix(camera.GetMatrixForPos(vec))
	}
	sprite.GetSprite().Draw(sp.batches[sprite], pixel.IM.Rotated(pixel.ZV, rotation).ScaledXY(pixel.ZV, scale).Moved(vec))
}

func (sp *SpriteBatch) End() error {
	if sp.state == 1 {
		for _, b := range sp.batches {
			b.Draw(sp.win.win)
		}
		for _, b := range sp.fontbatches {
			b.Draw(sp.win.win)
		}
		sp.state = 0
		return nil
	}
	return errors.New("Tried to end already ended batch!")
}
