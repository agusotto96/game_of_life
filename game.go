package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	bytesPerPixel = 4
	windowTitle   = "Game of Life"
)

type Color = [bytesPerPixel]byte

type Game struct {
	World        *World
	WindowWidth  int
	WindowHeight int
	TPS          int
	Pixels       []byte
	Alive        Color
	Dead         Color
}

var _ ebiten.Game = (*Game)(nil)

func NewColor(hexStr string) (Color, error) {
	hexStr = strings.TrimPrefix(hexStr, "#")

	switch len(hexStr) {
	case 6:
		hexStr += "ff"
	case 8:
	default:
		return Color{}, fmt.Errorf("invalid hexadecimal color %q: expected 6 (RRGGBB) or 8 (RRGGBBAA) characters", hexStr)
	}

	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return Color{}, fmt.Errorf("invalid hex color %q: %w", hexStr, err)
	}

	return Color(bytes), nil
}

func NewGame(width, height, cellSize, tps int, chance float64, rule string, alive, dead Color) (*Game, error) {
	if width <= 0 {
		return nil, errors.New("width must be greater than 0")
	}
	if height <= 0 {
		return nil, errors.New("height must be greater than 0")
	}
	if cellSize <= 0 {
		return nil, errors.New("cell size must be greater than 0")
	}
	if cellSize > width || cellSize > height {
		return nil, errors.New("cell size cannot be greater than window dimensions")
	}
	if tps <= 0 {
		return nil, errors.New("tps must be greater than 0")
	}

	gridWidth := width / cellSize
	gridHeight := height / cellSize
	world, err := NewWorld(gridWidth, gridHeight, chance, rule)
	if err != nil {
		return nil, err
	}

	pixels := make([]byte, world.Width*world.Height*bytesPerPixel)
	g := &Game{
		World:        world,
		WindowWidth:  width,
		WindowHeight: height,
		TPS:          tps,
		Pixels:       pixels,
		Alive:        alive,
		Dead:         dead,
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
	return g.World.Width, g.World.Height
}

func (g *Game) Run() error {
	ebiten.SetWindowTitle(windowTitle)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(g.WindowWidth, g.WindowHeight)
	ebiten.SetTPS(g.TPS)
	return ebiten.RunGame(g)
}
