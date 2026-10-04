package main

import (
	"math/rand"
)

const (
	neighborsToBirth   = 3
	neighborsToSurvive = 2
)

var neighborOffsets = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

type World struct {
	cells  []bool
	next   []bool
	Width  int
	Height int
}

func NewWorld(width int, height int, chance int) *World {
	cells := make([]bool, (width+2)*(height+2))
	for y := 1; y <= height; y++ {
		for x := 1; x <= width; x++ {
			cells[x+(y*(width+2))] = rand.Intn(chance) == 0
		}
	}
	next := make([]bool, len(cells))
	return &World{
		cells:  cells,
		next:   next,
		Width:  width,
		Height: height,
	}
}

func (w *World) At(x, y int) bool {
	return w.cells[(x+1)+(y+1)*(w.Width+2)]
}

func (w *World) Update() {
	for y := 1; y <= w.Height; y++ {
		for x := 1; x <= w.Width; x++ {
			i := x + (y * (w.Width + 2))
			n := w.aliveNeighbours(x, y)
			switch n {
			case neighborsToBirth:
				w.next[i] = true
			case neighborsToSurvive:
				w.next[i] = w.cells[i]
			default:
				w.next[i] = false
			}
		}
	}
	w.cells, w.next = w.next, w.cells
}

func (w *World) aliveNeighbours(x, y int) int {
	width := w.Width + 2
	count := 0
	for _, n := range neighborOffsets {
		x := x + n[0]
		y := y + n[1]
		if w.cells[x+y*width] {
			count++
		}
	}
	return count
}
