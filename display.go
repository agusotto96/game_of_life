package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const rgba = 4

type Color = [rgba]byte

type Game struct {
	Worlds <-chan World
	Width  int
	Height int
	TPS    int
	Pixels []byte
	Alive  Color
	Dead   Color
}

func NewGame(worlds <-chan World, width, height, tps int, alive, dead Color) *Game {
	g := &Game{
		Worlds: worlds,
		Width:  width,
		Height: height,
		TPS:    tps,
		Pixels: make([]byte, width*height*rgba),
		Alive:  alive,
		Dead:   dead,
	}
	return g
}

func RunGame(g *Game) error {
	ebiten.SetWindowTitle("Game of Life")
	if g.TPS > 0 {
		ebiten.SetTPS(g.TPS)
	}
	return ebiten.RunGame(g)
}

func (g *Game) updatePixels(w World) {
	for i, alive := range w.Cells {
		color := g.Dead
		if alive {
			color = g.Alive
		}
		for channel, value := range color {
			g.Pixels[i*rgba+channel] = value
		}
	}
}

func (g *Game) Update() error {
	w, ok := <-g.Worlds
	if !ok {
		return ebiten.Termination
	}
	g.updatePixels(w)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.Pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}
