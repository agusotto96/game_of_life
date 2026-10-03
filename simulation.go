package main

import (
	"context"
	"math/rand"
)

type World struct {
	Cells  []bool
	Width  int
	Height int
}

func RandomWorld(width int, height int, chance int) World {
	cells := make([]bool, width*height)
	for i := range cells {
		isAlive := rand.Intn(chance) == 0
		cells[i] = isAlive
	}
	return World{
		Cells:  cells,
		Width:  width,
		Height: height,
	}
}

func UpdateWorlds(ctx context.Context, world World) <-chan World {
	worlds := make(chan World)
	go func() {
		defer close(worlds)
		current := world.Cells
		next := make([]bool, len(current))
		width, height := world.Width, world.Height

		for {
			select {
			case <-ctx.Done():
				return
			case worlds <- World{Cells: current, Width: width, Height: height}:
			}
			stepWorld(current, next, width, height)
			current, next = next, current
		}
	}()
	return worlds
}

func updateWorld(w World) World {
	cells := make([]bool, len(w.Cells))
	stepWorld(w.Cells, cells, w.Width, w.Height)
	return World{
		Cells:  cells,
		Width:  w.Width,
		Height: w.Height,
	}
}

func stepWorld(src, dst []bool, width, height int) {
	w := World{Cells: src, Width: width, Height: height}
	for y := 0; y < height; y++ {
		rowOffset := y * width
		for x := 0; x < width; x++ {
			i := x + rowOffset
			n := aliveNeighbours(w, x, y)
			dst[i] = n == 3 || (n == 2 && src[i])
		}
	}
}

func aliveNeighbours(w World, x int, y int) int {
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
