package component

import (
	"github.com/hajimehoshi/ebiten"
)

type Bubble struct {
	Params        Object
	PositiveImage *ebiten.Image
	NegativeImage *ebiten.Image
	Positive      bool
}

func (b *Bubble) DrawOn(screen *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(1, 1)
	o.GeoM.Translate(b.Params.X, b.Params.Y)
	if b.Positive {
		screen.DrawImage(b.PositiveImage, o)
	} else {
		screen.DrawImage(b.NegativeImage, o)
	}
}

func (b *Bubble) Update(m *Magnet) {
	b.Params.Y += b.Params.VY
	if m.Positive {
		switch b.Positive {
		case true:
			b.Params.VY += 0.05
		case false:
			b.Params.VY -= 0.05
		}
	} else {
		switch b.Positive {
		case true:
			b.Params.VY -= 0.05
		case false:
			b.Params.VY += 0.05
		}
	}
	if b.Params.Y <= 0 || b.Params.Y >= ScreenHeight {
		b.Params.Die()
	}
}
