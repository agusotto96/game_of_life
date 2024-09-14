package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const rgba = 4

var aliveColor = [rgba]byte{0x39, 0xff, 0x14, 0xff}
var deadColor = [rgba]byte{0x20, 0x20, 0x20, 0xff}

type Game struct {
	World  World
	Pixels []byte
}

func NewGame(world World) Game {
	return Game{
		World:  world,
		Pixels: make([]byte, world.Width*world.Height*rgba),
	}
}

func (g *Game) Update() error {
	g.World = UpdateWorld(g.World)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i, alive := range g.World.Cells {
		color := deadColor
		if alive {
			color = aliveColor
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

func (g *Game) Run() error {
	ebiten.SetWindowTitle("Game of Life")
	return ebiten.RunGame(g)
}
