package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/theonlymoby/magnetib/assets"
	"github.com/theonlymoby/magnetib/component"
)

const (
	screenWidth  = 440
	screenHeight = 280
	tileSize     = 16
	objectSize   = tileSize * 2
)

type Mode int

const (
	ModeStart Mode = iota
	ModeGame
	ModeRetry
)

type Game struct {
	mode Mode

	hero   *component.Bubble
	magnet *component.Magnet
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func init() {
	assets.LoadStaticImages()

}

func (g *Game) Draw(screen *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(1, 1)
	screen.DrawImage(assets.BgImage, o)
	g.hero.DrawOn(screen)
	g.magnet.DrawOn(screen)
}

func (g *Game) InitObjects() {
	if g.hero == nil {
		g.hero = &component.Bubble{
			PositiveImage: assets.PositiveImage,
			NegativeImage: assets.NegativeImage,
			Width:         tileSize,
			Height:        tileSize,
			Positive:      true,
			Params: component.Object{
				X:      (screenWidth - tileSize) / 2 / tileSize,
				Y:      (screenHeight - tileSize) / 2,
				Alive:  true,
				CanDie: true,
			},
		}
	}
	if g.magnet == nil {
		g.magnet = &component.Magnet{
			MagnetPositiveImage: assets.MagnetPositiveImage,
			MagnetNegativeImage: assets.MagnetNegativeImage,
			Width:               tileSize,
			Height:              tileSize,
			Positive:            true,
			Params: component.Object{
				X:      (screenWidth - tileSize) / 2 / tileSize,
				Y:      (screenHeight - tileSize) / 2 / tileSize,
				Alive:  true,
				CanDie: false,
			},
		}
	}
}

func (g *Game) Update() error {

	g.InitObjects()

	switch g.mode {
	case ModeStart:
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.mode = ModeGame
		}
	case ModeGame:
		if ebiten.IsKeyPressed(ebiten.KeyJ) {
			g.hero.Positive = false
		}
		if ebiten.IsKeyPressed(ebiten.KeyK) {
			g.hero.Positive = true
		}
		g.hero.Update(g.magnet)
	}
	return nil
}

func main() {
	ebiten.RunGame(&Game{})
}
