package component

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/solarlune/goaseprite"
)

type Fly struct {
	Params    Object
	FlyImage  *ebiten.Image
	FlySprite *goaseprite.File
}

func (f *Fly) DrawOn(screen *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	sub := f.FlyImage.SubImage(image.Rect(f.FlySprite.CurrentFrameCoords()))
	o.GeoM.Scale(2, 2)
	o.GeoM.Translate(f.Params.X, f.Params.Y)
	screen.DrawImage(sub.(*ebiten.Image), o)
}

func (f *Fly) Update() {
	f.Params.X -= 1
}
