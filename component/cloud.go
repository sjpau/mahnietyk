package component

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/solarlune/goaseprite"
)

type Cloud struct {
	Params      Object
	CloudImage  *ebiten.Image
	CloudSprite *goaseprite.File
}

func (c *Cloud) DrawOn(screen *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	sub := c.CloudImage.SubImage(image.Rect(c.CloudSprite.CurrentFrameCoords()))
	o.GeoM.Scale(1, 1)
	o.GeoM.Translate(c.Params.X, c.Params.Y)
	screen.DrawImage(sub.(*ebiten.Image), o)
}

func (c *Cloud) Update() {
	c.Params.X -= 0.5
}
