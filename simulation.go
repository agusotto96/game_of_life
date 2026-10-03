package main

import (
	"math/rand"
)

type World struct {
	Cells  []bool
	next   []bool
	Width  int
	Height int
}

func NewWorld(width int, height int, chance int) *World {
	cells := make([]bool, width*height)
	for i := range cells {
		isAlive := rand.Intn(chance) == 0
		cells[i] = isAlive
	}
	next := make([]bool, width*height)
	return &World{
		Cells:  cells,
		next:   next,
		Width:  width,
		Height: height,
	}
}

func (w *World) Update() {
	for y := range w.Height {
		for x := range w.Width {
			i := x + (y * w.Width)
			n := w.aliveNeighbours(x, y)
			switch n {
			case 3:
				w.next[i] = true
			case 2:
				w.next[i] = w.Cells[i]
			default:
				w.next[i] = false
			}
		}
	}
	w.Cells, w.next = w.next, w.Cells
}

func (w *World) aliveNeighbours(x, y int) int {
	count := 0
	neighbors := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, n := range neighbors {
		x := x + n[0]
		y := y + n[1]
		if x < 0 || y < 0 || x >= w.Width || y >= w.Height {
			continue
		}
		isAlive := w.Cells[x+y*w.Width]
		if isAlive {
			count++
		}
	}
	return count
}
