package main

import (
	"flag"
	"log"
)

const (
	defaultWidth    = 960
	defaultHeight   = 600
	defaultCellSize = 2
	defaultChance   = 0.15
	defaultTPS      = 60
	defaultRule     = "B3/S23"
	defaultAliveHex = "39ff14ff"
	defaultDeadHex  = "202020ff"
)

func main() {
	width := flag.Int("width", defaultWidth, "The window width in pixels.")
	height := flag.Int("height", defaultHeight, "The window height in pixels.")
	cellSize := flag.Int("cell-size", defaultCellSize, "The display size of each cell in pixels.")
	chance := flag.Float64("chance", defaultChance, "The probability (0.0 to 1.0) that a cell starts alive.")
	tps := flag.Int("tps", defaultTPS, "The number of simulation updates per second (Ticks Per Second).")
	rule := flag.String("rule", defaultRule, "Life-like cellular automaton rule in B.../S... format (e.g. B3/S23, B36/S23, B35678/S5678).")
	aliveHex := flag.String("alive", defaultAliveHex, "Hexadecimal RGBA color for alive cells.")
	deadHex := flag.String("dead", defaultDeadHex, "Hexadecimal RGBA color for dead cells.")
	flag.Parse()

	alive, err := NewColor(*aliveHex)
	if err != nil {
		log.Fatal(err)
	}
	dead, err := NewColor(*deadHex)
	if err != nil {
		log.Fatal(err)
	}

	game, err := NewGame(*width, *height, *cellSize, *tps, *chance, *rule, alive, dead)
	if err != nil {
		log.Fatal(err)
	}

	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
