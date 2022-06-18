package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/theonlymoby/magnetib/assets"
	"github.com/theonlymoby/magnetib/component"
)

type Mode int

const (
	ModeStart Mode = iota
	ModeGame
	ModeRetry
)

type Game struct {
	mode Mode

	bubble *component.Bubble
	magnet *component.Magnet
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return component.ScreenWidth, component.ScreenHeight
}

func init() {
	assets.LoadStaticImages()

}

func (g *Game) Draw(screen *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(1, 1)
	screen.DrawImage(assets.BgImage, o)
	g.bubble.DrawOn(screen)
	g.magnet.DrawOn(screen)
}

func (g *Game) InitObjects() {

	if g.bubble == nil {
		g.bubble = &component.Bubble{
			PositiveImage: assets.PositiveImage,
			NegativeImage: assets.NegativeImage,
			Positive:      true,
			Params: component.Object{
				X:      (component.ScreenWidth - component.TileSize) / 2 / component.TileSize,
				Y:      (component.ScreenHeight - component.TileSize) / 2,
				Alive:  true,
				CanDie: true,
				Width:  component.TileSize,
				Height: component.TileSize,
			},
		}
	}
	if g.magnet == nil {
		g.magnet = &component.Magnet{
			MagnetPositiveImage: assets.MagnetPositiveImage,
			MagnetNegativeImage: assets.MagnetNegativeImage,
			Positive:            true,
			Params: component.Object{
				X:      (component.ScreenWidth - component.TileSize) / 2 / component.TileSize,
				Y:      (component.ScreenHeight - component.TileSize) / 2 / component.TileSize,
				Alive:  true,
				CanDie: false,
				Width:  component.TileSize,
				Height: component.TileSize,
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
			g.bubble.Positive = false
		}
		if ebiten.IsKeyPressed(ebiten.KeyK) {
			g.bubble.Positive = true
		}
		if g.bubble.Params.Alive == false {
			g.mode = ModeRetry
		}
		if g.bubble.Params.CollideWith(&g.magnet.Params) {
			g.bubble.Die()
		}
		g.bubble.Update(g.magnet)
	case ModeRetry:
		g.bubble = nil
		g.magnet = nil
		g.InitObjects()
		g.mode = ModeStart
	}
	return nil
}

func main() {
	ebiten.RunGame(&Game{})
}
