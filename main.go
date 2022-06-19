package main

import (
	_ "image/png"
	"math/rand"
	"time"

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

type Flies struct {
	flies []*component.Fly
	spawn int
}

func (f *Flies) Update() {
	for i := 0; i < f.spawn; i++ {
		f.flies[i].Update()
	}
}

func (f *Flies) DrawOn(screen *ebiten.Image) {
	for i := 0; i < f.spawn; i++ {
		f.flies[i].DrawOn(screen)
	}
}

type Game struct {
	mode Mode

	bubble *component.Bubble
	magnet *component.Magnet
	flies  Flies
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return component.ScreenWidth, component.ScreenHeight
}

func init() {
	rand.Seed(time.Now().UnixNano())
	assets.LoadStaticImages()
	assets.LoadDynamicImages()
}

func (g *Game) Draw(screen *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(1, 1)
	screen.DrawImage(assets.BgImage, o)
	g.bubble.DrawOn(screen)
	g.magnet.DrawOn(screen)
	g.flies.DrawOn(screen)
}

func (g *Game) InitObjects() {

	if g.flies.flies == nil {
		g.flies.flies = make([]*component.Fly, 19)
		g.flies.spawn = 1
		for i := range g.flies.flies {
			seed := rand.Intn(17)
			g.flies.flies[i] = &component.Fly{
				FlySprite: assets.FlySprite,
				FlyImage:  assets.FlyImage,
				Params: component.Object{
					X:      component.ScreenWidth - component.TileSize,
					Y:      float64(component.ScreenHeight - (component.TileSize * seed)),
					VX:     0,
					VY:     0,
					Alive:  true,
					Width:  16,
					Height: 16,
				},
			}
		}
	}

	if g.bubble == nil {
		g.bubble = new(component.Bubble)
		g.bubble = &component.Bubble{
			PositiveImage: assets.PositiveImage,
			NegativeImage: assets.NegativeImage,
			Positive:      true,
			Params: component.Object{
				X:      (component.ScreenWidth - component.TileSize) / 2 / component.TileSize,
				Y:      (component.ScreenHeight - component.TileSize) / 2,
				Alive:  true,
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
		assets.FlySprite.Play("run")
		assets.FlySprite.Update(float32(1.0 / 60.0))
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
			g.bubble.Params.Die()
		}
		g.bubble.Update(g.magnet)
		g.flies.Update()
	case ModeRetry:
		g.bubble = nil
		g.flies.flies = nil
		g.magnet = nil
		g.InitObjects()
		g.mode = ModeStart
	}
	return nil
}

func main() {
	ebiten.RunGame(&Game{})
}
