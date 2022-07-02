package assets

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

var (
	BgImage        *ebiten.Image
	MenuImage      *ebiten.Image
	PositiveImage  *ebiten.Image
	NegativeImage  *ebiten.Image
	PositiveSprite *goaseprite.File
	NegativeSprite *goaseprite.File
	MagnetImage    *ebiten.Image
	MagnetSprite   *goaseprite.File
	FlySprite      *goaseprite.File
	FlyImage       *ebiten.Image
	CloudSprite    *goaseprite.File
	CloudImage     *ebiten.Image
)

var (
	AudioContext *audio.Context
	BubbleCharge *audio.Player
	BubbleHit    *audio.Player
	MagnetCharge *audio.Player
	FlyGull      *audio.Player
	CloudThunder *audio.Player
)

const (
	SampleRate = 48000
)

func LoadAudio() {
	AudioContext = audio.NewContext(SampleRate)

	temp, err := wav.DecodeWithSampleRate(SampleRate, bytes.NewReader(bubble_charge_wav))
	if err != nil {
		panic(err)
	}
	BubbleCharge, err = AudioContext.NewPlayer(temp)
	if err != nil {
		panic(err)
	}
	temp, err = wav.DecodeWithSampleRate(SampleRate, bytes.NewReader(magnet_charge_wav))
	if err != nil {
		panic(err)
	}
	MagnetCharge, err = AudioContext.NewPlayer(temp)
	if err != nil {
		panic(err)
	}

	temp, err = wav.DecodeWithSampleRate(SampleRate, bytes.NewReader(bubble_hit_wav))
	if err != nil {
		panic(err)
	}
	BubbleHit, err = AudioContext.NewPlayer(temp)
	if err != nil {
		panic(err)
	}
	temp, err = wav.DecodeWithSampleRate(SampleRate, bytes.NewReader(gull_wav))
	if err != nil {
		panic(err)
	}
	FlyGull, err = AudioContext.NewPlayer(temp)
	if err != nil {
		panic(err)
	}
	temp, err = wav.DecodeWithSampleRate(SampleRate, bytes.NewReader(thunder_wav))
	if err != nil {
		panic(err)
	}
	CloudThunder, err = AudioContext.NewPlayer(temp)
	if err != nil {
		panic(err)
	}
}

func LoadDynamicImages() {
	FlySprite = goaseprite.Open("assets/img/fly.json")
	img, _, err := ebitenutil.NewImageFromFile(FlySprite.ImagePath)
	if err != nil {
		panic(err)
	}
	FlyImage = img
	FlySprite.Play("run")

	MagnetSprite = goaseprite.Open("assets/img/magnet.json")
	img, _, err = ebitenutil.NewImageFromFile(MagnetSprite.ImagePath)
	if err != nil {
		panic(err)
	}
	MagnetImage = img

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

	img, _, err = image.Decode(bytes.NewReader(menu_png))
	if err != nil {
		panic(err)
	}
	MenuImage = ebiten.NewImageFromImage(img)
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
	MagnetSprite.Update(float32(1.0 / 60.0))

}
