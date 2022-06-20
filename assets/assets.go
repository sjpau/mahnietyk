package assets

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/solarlune/goaseprite"
)

var (
	BgImage             *ebiten.Image
	PositiveImage       *ebiten.Image
	NegativeImage       *ebiten.Image
	MagnetPositiveImage *ebiten.Image
	MagnetNegativeImage *ebiten.Image
	FlySprite           *goaseprite.File
	FlyImage            *ebiten.Image
	CloudSprite         *goaseprite.File
	CloudImage          *ebiten.Image
)

func LoadDynamicImages() {
	FlySprite = goaseprite.Open("assets/img/fly.json")
	img, _, err := ebitenutil.NewImageFromFile(FlySprite.ImagePath)
	if err != nil {
		panic(err)
	}
	FlyImage = img
	FlySprite.Play("run")

	CloudSprite = goaseprite.Open("assets/img/cloud.json")
	img, _, err = ebitenutil.NewImageFromFile(CloudSprite.ImagePath)
	if err != nil {
		panic(err)
	}
	CloudImage = img
	CloudSprite.Play("run")
}

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
