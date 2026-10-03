package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const rgba = 4

type Color = [rgba]byte

type Game struct {
	World  *World
	Width  int
	Height int
	TPS    int
	Pixels []byte
	Alive  Color
	Dead   Color
}

func NewGame(world *World, tps int, alive, dead Color) *Game {
	pixels := make([]byte, world.Width*world.Height*rgba)
	g := &Game{
		World:  world,
		Width:  world.Width,
		Height: world.Height,
		TPS:    tps,
		Pixels: pixels,
		Alive:  alive,
		Dead:   dead,
	}
	return g
}

func RunGame(g *Game) error {
	ebiten.SetWindowTitle("Game of Life")
	ebiten.SetTPS(g.TPS)
	return ebiten.RunGame(g)
}

func (g *Game) Update() error {
	for i, alive := range g.World.Cells {
		color := g.Dead
		if alive {
			color = g.Alive
		}
		for channel, value := range color {
			g.Pixels[i*rgba+channel] = value
		}
	}
	g.World.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.Pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}
