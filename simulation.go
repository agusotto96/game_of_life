package main

import "math/rand"

type World struct {
	Cells  []bool
	Width  int
	Height int
}

func NewWorld(width int, height int, chance int) World {
	cells := make([]bool, width*height)
	for i := 0; i < len(cells); i++ {
		cells[i] = rand.Intn(chance) == 0
	}
	return World{
		Cells:  cells,
		Width:  width,
		Height: height,
	}
}

func (w *World) Update() {
	cells := make([]bool, w.Width*w.Height)
	for x := 0; x < w.Width; x++ {
		for y := 0; y < w.Height; y++ {
			i := x + y*w.Width
			n := w.aliveNeighbours(x, y)
			switch {
			case n == 2 && w.Cells[i]:
				cells[i] = true
			case n == 3:
				cells[i] = true
			default:
				cells[i] = false
			}
		}
	}
	w.Cells = cells
}

func (w *World) aliveNeighbours(x int, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			x := x + j
			y := y + i
			if x < 0 || y < 0 || x >= w.Width || y >= w.Height {
				continue
			}
			if w.Cells[x+y*w.Width] {
				count++
			}
		}
	}
	return count
}
