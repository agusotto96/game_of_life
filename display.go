package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const rgba = 4

type Color = [rgba]byte

type Game struct {
	World  World
	Pixels []byte
	Alive  Color
	Dead   Color
}

func NewGame(world World, alive, dead Color) Game {
	return Game{
		World:  world,
		Pixels: make([]byte, world.Width*world.Height*rgba),
		Alive:  alive,
		Dead:   dead,
	}
}

func RunGame(g Game) error {
	ebiten.SetWindowTitle("Game of Life")
	return ebiten.RunGame(&g)
}

func (g *Game) Update() error {
	g.World = UpdateWorld(g.World)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i, alive := range g.World.Cells {
		color := g.Dead
		if alive {
			color = g.Alive
		}
		for channel, value := range color {
			g.Pixels[i*rgba+channel] = value
		}
	}
	screen.WritePixels(g.Pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.World.Width, g.World.Height
}
