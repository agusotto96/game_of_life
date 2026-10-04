package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	bytesPerPixel = 4
	windowTitle   = "Game of Life"
)

type Color = [bytesPerPixel]byte

type Game struct {
	World  *World
	Width  int
	Height int
	TPS    int
	Pixels []byte
	Alive  Color
	Dead   Color
}

var _ ebiten.Game = (*Game)(nil)

func NewGame(world *World, tps int, alive, dead Color) *Game {
	pixels := make([]byte, world.Width*world.Height*bytesPerPixel)
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
	ebiten.SetWindowTitle(windowTitle)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(g.Width, g.Height)
	ebiten.SetTPS(g.TPS)
	return ebiten.RunGame(g)
}

func (g *Game) Update() error {
	g.World.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	pixIdx := 0
	for y := range g.Height {
		for x := range g.Width {
			color := g.Dead
			if g.World.At(x, y) {
				color = g.Alive
			}
			for channel, value := range color {
				g.Pixels[pixIdx+channel] = value
			}
			pixIdx += bytesPerPixel
		}
	}
	screen.WritePixels(g.Pixels)
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}
