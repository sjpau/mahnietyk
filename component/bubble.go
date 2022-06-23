package component

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/solarlune/goaseprite"
)

type Bubble struct {
	Params         Object
	PositiveImage  *ebiten.Image
	NegativeImage  *ebiten.Image
	PositiveSprite *goaseprite.File
	NegativeSprite *goaseprite.File
	Positive       bool
}

func (b *Bubble) DrawOn(screen *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(1, 1)
	o.GeoM.Translate(b.Params.X, b.Params.Y)
	subp := b.PositiveImage.SubImage(image.Rect(b.PositiveSprite.CurrentFrameCoords()))
	subn := b.NegativeImage.SubImage(image.Rect(b.NegativeSprite.CurrentFrameCoords()))
	if b.Positive {
		screen.DrawImage(subp.(*ebiten.Image), o)
	} else {
		screen.DrawImage(subn.(*ebiten.Image), o)
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

func (b *Bubble) ChangeCharge() {
	if b.Positive {
		b.Positive = false
	} else {
		b.Positive = true
	}
}
