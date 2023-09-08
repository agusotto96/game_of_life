package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Width  int
	Height int
	World  World
	Pixels []byte
}

func NewGame(world World) Game {
	return Game{
		Width:  world.Width,
		Height: world.Height,
		World:  world,
		Pixels: make([]byte, world.Width*world.Height*4),
	}
}

func (g *Game) Update() error {
	g.World.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i, v := range g.World.Cells {
		if v {
			g.Pixels[4*i] = 0xff
			g.Pixels[4*i+1] = 0xff
			g.Pixels[4*i+2] = 0xff
			g.Pixels[4*i+3] = 0xff
		} else {
			g.Pixels[4*i] = 0
			g.Pixels[4*i+1] = 0
			g.Pixels[4*i+2] = 0
			g.Pixels[4*i+3] = 0
		}
	}
	screen.WritePixels(g.Pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}

func (g *Game) Run() error {
	return ebiten.RunGame(g)
}
