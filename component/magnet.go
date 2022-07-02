package component

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/goaseprite"
)

type Magnet struct {
	Params       Object
	MagnetImage  *ebiten.Image
	MagnetSprite *goaseprite.File
	Positive     bool
}

func (m *Magnet) DrawOn(screen *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(1, -1)
	o.GeoM.Translate(0, 16)
	o.GeoM.Translate(m.Params.X, m.Params.Y)
	sub := m.MagnetImage.SubImage(image.Rect(m.MagnetSprite.CurrentFrameCoords()))
	if m.Positive {
		screen.DrawImage(sub.(*ebiten.Image), o)
		m.MagnetSprite.Play("positive")
	} else {
		screen.DrawImage(sub.(*ebiten.Image), o)
		m.MagnetSprite.Play("negative")
	}
}
