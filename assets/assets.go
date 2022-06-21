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
	PositiveSprite      *goaseprite.File
	NegativeSprite      *goaseprite.File
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

	PositiveSprite = goaseprite.Open("assets/img/positive.json")
	img, _, err = ebitenutil.NewImageFromFile(PositiveSprite.ImagePath)
	if err != nil {
		panic(err)
	}
	PositiveImage = img
	PositiveSprite.Play("run")

	NegativeSprite = goaseprite.Open("assets/img/negative.json")
	img, _, err = ebitenutil.NewImageFromFile(NegativeSprite.ImagePath)
	if err != nil {
		panic(err)
	}
	NegativeImage = img
	NegativeSprite.Play("run")
}

func LoadStaticImages() {
	img, _, err := image.Decode(bytes.NewReader(background_png))
	if err != nil {
		panic(err)
	}
	BgImage = ebiten.NewImageFromImage(img)

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

func PlayAssets() {
	FlySprite.Play("run")
	FlySprite.Update(float32(1.0 / 60.0))
	CloudSprite.Play("run")
	CloudSprite.Update(float32(1.0 / 60.0))
	PositiveSprite.Play("run")
	PositiveSprite.Update(float32(1.0 / 60.0))
	NegativeSprite.Play("run")
	NegativeSprite.Update(float32(1.0 / 60.0))

}
