package main

import (
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

func UpdateWorlds(world World) <-chan World {
	worlds := make(chan World)
	go func() {
		for {
			world = updateWorld(world)
			worlds <- world
		}
	}()
	return worlds
}

func updateWorld(w World) World {
	cells := make([]bool, w.Width*w.Height)
	for x := 0; x < w.Width; x++ {
		for y := 0; y < w.Height; y++ {
			i := x + y*w.Width
			n := aliveNeighbours(w, x, y)
			isAlive := n == 3 || n == 2 && w.Cells[i]
			cells[i] = isAlive
		}
	}
	return World{
		Cells:  cells,
		Width:  w.Width,
		Height: w.Height,
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
