package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const rgba = 4

type Color = [rgba]byte

type Game struct {
	Worlds chan World
	Width  int
	Height int
	Pixels []byte
	Alive  Color
	Dead   Color
}

func NewGame(world World, alive, dead Color) Game {
	g := Game{
		Worlds: make(chan World, 1000),
		Width:  world.Width,
		Height: world.Height,
		Pixels: make([]byte, world.Width*world.Height*rgba),
		Alive:  alive,
		Dead:   dead,
	}
	go func() {
		for {
			world = UpdateWorld(world)
			g.Worlds <- world
		}
	}()
	return g
}

func RunGame(g Game) error {
	ebiten.SetWindowTitle("Game of Life")
	return ebiten.RunGame(&g)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w := <-g.Worlds
	for i, alive := range w.Cells {
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
	return g.Width, g.Height
}
