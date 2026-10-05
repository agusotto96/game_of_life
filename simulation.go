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

type Rule struct {
	raw     string
	birth   [maxNeighbors + 1]bool
	survive [maxNeighbors + 1]bool
}

func NewRule(s string) (Rule, error) {
	matches := ruleRegex.FindStringSubmatch(strings.ToUpper(strings.TrimSpace(s)))
	if len(matches) != 3 {
		return Rule{}, errors.New("invalid rule format, expected B.../S... (e.g. B3/S23)")
	}

	var rule Rule
	rule.raw = matches[0]

	for _, ch := range matches[1] {
		n, _ := strconv.Atoi(string(ch))
		rule.birth[n] = true
	}
	for _, ch := range matches[2] {
		n, _ := strconv.Atoi(string(ch))
		rule.survive[n] = true
	}

	return rule, nil
}

func (r *Rule) String() string {
	return r.raw
}

func (r *Rule) nextState(alive bool, neighbors int) bool {
	if alive {
		return r.survive[neighbors]
	}
	return r.birth[neighbors]
}

type World struct {
	Cells  []bool
	next   []bool
	Width  int
	Height int
	Rule   Rule
}

func NewWorld(width int, height int, chance float64, rule Rule) *World {
	cells := make([]bool, width*height)
	for i := range cells {
		cells[i] = rand.Float64() < chance
	}
	next := make([]bool, width*height)
	return &World{
		Cells:  cells,
		next:   next,
		Width:  width,
		Height: height,
		Rule:   rule,
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
					w.next[i] = w.Rule.nextState(w.Cells[i], n)
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
