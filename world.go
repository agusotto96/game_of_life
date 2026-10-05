package main

import (
	"errors"
	"math/rand"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

const maxNeighbors = 8

var (
	ruleRegex       = regexp.MustCompile(`^B([0-8]*)/S([0-8]*)$`)
	neighborOffsets = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
)

type World struct {
	Cells   []bool
	next    []bool
	Width   int
	Height  int
	Rule    string
	birth   [maxNeighbors + 1]bool
	survive [maxNeighbors + 1]bool
}

func NewWorld(width int, height int, chance float64, rule string) (*World, error) {
	if width <= 0 {
		return nil, errors.New("width must be greater than 0")
	}
	if height <= 0 {
		return nil, errors.New("height must be greater than 0")
	}
	if chance <= 0 || chance > 1 {
		return nil, errors.New("chance must be between 0.0 and 1.0")
	}

	matches := ruleRegex.FindStringSubmatch(strings.ToUpper(strings.TrimSpace(rule)))
	if len(matches) != 3 {
		return nil, errors.New("invalid rule format, expected B.../S... (e.g. B3/S23)")
	}

	var birth, survive [maxNeighbors + 1]bool
	for _, ch := range matches[1] {
		n, _ := strconv.Atoi(string(ch))
		birth[n] = true
	}
	for _, ch := range matches[2] {
		n, _ := strconv.Atoi(string(ch))
		survive[n] = true
	}

	cells := make([]bool, width*height)
	for i := range cells {
		cells[i] = rand.Float64() < chance
	}
	next := make([]bool, width*height)

	return &World{
		Cells:   cells,
		next:    next,
		Width:   width,
		Height:  height,
		Rule:    matches[0],
		birth:   birth,
		survive: survive,
	}, nil
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
					w.next[i] = w.nextState(w.Cells[i], n)
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

func (w *World) nextState(alive bool, neighbors int) bool {
	if alive {
		return w.survive[neighbors]
	}
	return w.birth[neighbors]
}
