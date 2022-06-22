package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/theonlymoby/mahnietyk/assets"
	"github.com/theonlymoby/mahnietyk/component"
)

type Mode int

const (
	ModeStart Mode = iota
	ModeGame
	ModeRetry
)

const (
	maxFlies  = 256
	maxClouds = 256
)

type Flies struct {
	flies []*component.Fly
	spawn int
}

type Clouds struct {
	clouds []*component.Cloud
	spawn  int
}

func (c *Clouds) Update(g *Game) {
	for i := 0; i < c.spawn; i++ {
		c.clouds[i].Update()
		if g.bubble.Params.CollideWith(&c.clouds[i].Params) {
			g.bubble.ChangeCharge()
		}
	}
}

func (c *Clouds) DrawOn(screen *ebiten.Image) {
	for i := 0; i < c.spawn; i++ {
		c.clouds[i].DrawOn(screen)
	}
}

func (f *Flies) Update(g *Game) {

	for i := 0; i < f.spawn; i++ {
		f.flies[i].Update()
		if g.bubble.Params.CollideWith(&f.flies[i].Params) {
			g.bubble.Params.Die()
		}
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
	clouds Clouds
	score  uint64
	points float64
	hard   bool
}

func (g *Game) EventMagnetChangeCharge() {
	if g.score%uint64((100+rand.Intn(1500))) == 0 {
		if g.magnet.Positive {
			g.magnet.Positive = false
		} else {
			g.magnet.Positive = true
		}
	}
}

func (g *Game) EventSpawnCloud() {
	if g.score%uint64(100+rand.Intn(1000)) == 0 && g.clouds.spawn < maxClouds {
		g.clouds.spawn += 1
	}
}

func (g *Game) EventSpawnFly() {
	if g.score%uint64(100+rand.Intn(20)) == 0 && g.flies.spawn < maxFlies {
		g.flies.spawn += 1
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return component.ScreenWidth, component.ScreenHeight
}

func init() {
	rand.Seed(time.Now().UnixNano())
	assets.LoadStaticImages()
	assets.LoadDynamicImages()
	assets.LoadFonts()
}

func (g *Game) Draw(screen *ebiten.Image) {

	var title []string
	var output []string

	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(1, 1)
	screen.DrawImage(assets.BgImage, o)
	pointString := fmt.Sprintf("%04f", g.points)
	switch g.mode {
	case ModeStart:
		title = []string{"MAHNIETYK"}
		output = []string{"", "", "J - gain negative charge", "", "K - gain positive charge", "", "", "Press SPACE to start"}
		for i, l := range title {
			x := (component.ScreenWidth - len(l)*assets.FontSize) / 2
			text.Draw(screen, l, assets.GameFont, x, (i+4)*assets.FontSize, color.White)
		}
		for i, l := range output {
			x := (component.ScreenWidth-len(l)*assets.FontSize)/2 + 2*component.TileSize
			text.Draw(screen, l, assets.GameFont, x, (i+4)*assets.FontSize, color.White)
		}
	case ModeGame:
		g.bubble.DrawOn(screen)
		g.magnet.DrawOn(screen)
		g.flies.DrawOn(screen)
		g.clouds.DrawOn(screen)
		text.Draw(screen, pointString, assets.GameFont, component.ScreenWidth-len(pointString)*assets.FontSize, assets.FontSize, color.White)
	case ModeRetry:
		//retry whatever goes here
	}
}

func (g *Game) InitObjects() {

	if g.flies.flies == nil {
		g.flies.flies = make([]*component.Fly, maxFlies)
		g.flies.spawn = 0
		for i := range g.flies.flies {
			seed := rand.Intn(17)
			g.flies.flies[i] = &component.Fly{
				FlySprite: assets.FlySprite,
				FlyImage:  assets.FlyImage,
				Params: component.Object{
					X:      component.ScreenWidth + component.TileSize,
					Y:      float64(component.ScreenHeight - (component.TileSize * seed)),
					VX:     0,
					VY:     0,
					Alive:  true,
					Width:  component.TileSize * 1.5,
					Height: component.TileSize * 1.5,
				},
			}
		}
	}

	if g.clouds.clouds == nil {
		g.clouds.clouds = make([]*component.Cloud, maxClouds)
		g.clouds.spawn = 0
		for i := range g.clouds.clouds {
			seed := rand.Intn(17)
			g.clouds.clouds[i] = &component.Cloud{
				CloudSprite: assets.CloudSprite,
				CloudImage:  assets.CloudImage,
				Params: component.Object{
					X:      component.ScreenWidth + component.TileSize,
					Y:      float64(component.ScreenHeight - (component.TileSize * seed)),
					VX:     0,
					VY:     0,
					Alive:  true,
					Width:  component.TileSize * 1.5,
					Height: component.TileSize * 1.5,
				},
			}
		}
	}

	if g.bubble == nil {
		g.bubble = new(component.Bubble)
		g.bubble = &component.Bubble{
			PositiveImage:  assets.PositiveImage,
			PositiveSprite: assets.PositiveSprite,
			NegativeImage:  assets.NegativeImage,
			NegativeSprite: assets.NegativeSprite,
			Positive:       true,
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
			MagnetImage:  assets.MagnetImage,
			MagnetSprite: assets.MagnetSprite,
			Positive:     true,
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
		assets.PlayAssets()
		if ebiten.IsKeyPressed(ebiten.KeyJ) {
			g.bubble.Positive = false
		}
		if ebiten.IsKeyPressed(ebiten.KeyK) {
			g.bubble.Positive = true
		}
		if g.bubble.Params.CollideWith(&g.magnet.Params) {
			g.bubble.Params.Die()
		}
		if g.bubble.Params.Alive == false {
			g.mode = ModeRetry
		}
		g.EventSpawnFly()
		g.flies.Update(g)
		g.EventSpawnCloud()
		g.clouds.Update(g)
		g.EventMagnetChangeCharge()
		g.bubble.Update(g.magnet)
		g.score += 1
		g.points += 0.000001
	case ModeRetry:
		g.bubble = nil
		g.flies.flies = nil
		g.clouds.clouds = nil
		g.magnet = nil
		g.flies.spawn = 0
		g.clouds.spawn = 0
		g.score = 0
		g.points = 0
		g.InitObjects()
		g.mode = ModeStart
	}
	return nil
}

func main() {
	ebiten.RunGame(&Game{})
}
