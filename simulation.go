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

func UpdateWorld(w World) World {
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
			isAlive := w.Cells[x+y*w.Width]
			if isAlive {
				count++
			}
		}
	}
	return count
}
