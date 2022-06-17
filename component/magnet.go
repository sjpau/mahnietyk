package component

import "github.com/hajimehoshi/ebiten"

type Magnet struct {
	Params              Object
	MagnetPositiveImage *ebiten.Image
	MagnetNegativeImage *ebiten.Image
	Positive            bool
	Width               int
	Height              int
}

func (m *Magnet) DrawOn(screen *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(1, -1)
	o.GeoM.Translate(0, 16)
	o.GeoM.Translate(m.Params.X, m.Params.Y)
	if m.Positive {
		screen.DrawImage(m.MagnetPositiveImage, o)
	} else {
		screen.DrawImage(m.MagnetNegativeImage, o)
	}
}
