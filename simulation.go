package main

import (
	"math/rand"
	"runtime"
	"sync"
)

const (
	neighborsToBirth   = 3
	neighborsToSurvive = 2
)

var neighborOffsets = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

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
	numWorkers := runtime.NumCPU()
	if numWorkers > w.Height {
		numWorkers = w.Height
	}

	chunkSize := w.Height / numWorkers
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for worker := range numWorkers {
		startY := worker * chunkSize
		endY := startY + chunkSize
		if worker == numWorkers-1 {
			endY = w.Height
		}

		go func(fromY, toY int) {
			defer wg.Done()
			for y := fromY; y < toY; y++ {
				for x := range w.Width {
					i := x + y*w.Width
					n := w.aliveNeighbours(x, y)
					switch n {
					case neighborsToBirth:
						w.next[i] = true
					case neighborsToSurvive:
						w.next[i] = w.Cells[i]
					default:
						w.next[i] = false
					}
				}
			}
		}(startY, endY)
	}

	wg.Wait()
	w.Cells, w.next = w.next, w.Cells
}

func (w *World) aliveNeighbours(x, y int) int {
	count := 0
	for _, n := range neighborOffsets {
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
