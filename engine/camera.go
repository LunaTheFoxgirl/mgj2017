package engine

import (
	"github.com/faiface/pixel"
)

type Camera struct {
	matrix pixel.Matrix
	trpos pixel.Vec
	trrot float64
	trscal float64
}

func (cam *Camera) GetMatrixForPos(pos pixel.Vec) pixel.Matrix {
	return cam.matrix.Moved(pos)
}

func (cam *Camera) Move(vec pixel.Vec) {
	cam.trpos = vec
}

func (cam *Camera) Rotate(angle float64) {
	cam.trrot = angle
}

func (cam *Camera) Scale(scale float64) {
	cam.trscal = scale
}

func (cam *Camera) Update() {
	cam.matrix.Rotated(pixel.ZV, cam.trrot)
	cam.matrix.Scaled(pixel.ZV, cam.trscal)
	cam.matrix.Moved(cam.trpos)
}