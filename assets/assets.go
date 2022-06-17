package assets

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

var FS embed.FS

var (
	BgImage             *ebiten.Image
	PositiveImage       *ebiten.Image
	NegativeImage       *ebiten.Image
	MagnetPositiveImage *ebiten.Image
	MagnetNegativeImage *ebiten.Image
)

func LoadStaticImages() {
	img, _, err := image.Decode(bytes.NewReader(background_png))
	if err != nil {
		panic(err)
	}
	BgImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(negative_png))
	if err != nil {
		panic(err)
	}
	NegativeImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(positive_png))
	if err != nil {
		panic(err)
	}
	PositiveImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(magnetpositive_png))
	if err != nil {
		panic(err)
	}
	MagnetPositiveImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(magnetnegative_png))
	if err != nil {
		panic(err)
	}
	MagnetNegativeImage = ebiten.NewImageFromImage(img)

}
