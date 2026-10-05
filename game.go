package main

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	bytesPerPixel  = 4
	hexColorLength = 8
	windowTitle    = "Game of Life"
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

func NewColor(hexStr string) (Color, error) {
	if len(hexStr) != hexColorLength {
		return Color{}, fmt.Errorf("invalid hexadecimal color format, expected %d characters (RRGGBBAA)", hexColorLength)
	}
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return Color{}, err
	}
	return (Color)(bytes), nil
}

func NewGame(world *World, tps int, alive, dead Color) (*Game, error) {
	if world == nil {
		return nil, errors.New("world cannot be nil")
	}
	if tps <= 0 {
		return nil, errors.New("tps must be greater than 0")
	}

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
	return g, nil
}

func (g *Game) Update() error {
	g.World.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i, alive := range g.World.Cells {
		color := g.Dead
		if alive {
			color = g.Alive
		}
		for channel, value := range color {
			g.Pixels[i*bytesPerPixel+channel] = value
		}
	}
	screen.WritePixels(g.Pixels)
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}

func (g *Game) Run() error {
	ebiten.SetWindowTitle(windowTitle)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(g.Width, g.Height)
	ebiten.SetTPS(g.TPS)
	return ebiten.RunGame(g)
}
