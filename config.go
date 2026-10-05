package main

import (
	"errors"
	"flag"
)

const (
	defaultWidth    = 960
	defaultHeight   = 600
	defaultChance   = 0.15
	defaultTPS      = 60
	defaultRule     = "B3/S23"
	defaultAliveHex = "39ff14ff"
	defaultDeadHex  = "202020ff"
)

type Config struct {
	Width  int
	Height int
	Chance float64
	TPS    int
	Rule   Rule
	Alive  Color
	Dead   Color
}

func NewConfig(width, height int, chance float64, tps int, rule Rule, alive, dead Color) (Config, error) {
	if width <= 0 {
		return Config{}, errors.New("width must be greater than 0")
	}
	if height <= 0 {
		return Config{}, errors.New("height must be greater than 0")
	}
	if chance <= 0 || chance > 1 {
		return Config{}, errors.New("chance must be between 0.0 and 1.0")
	}
	if tps <= 0 {
		return Config{}, errors.New("tps must be greater than 0")
	}
	return Config{
		Width:  width,
		Height: height,
		Chance: chance,
		TPS:    tps,
		Rule:   rule,
		Alive:  alive,
		Dead:   dead,
	}, nil
}

func ReadConfig() (Config, error) {
	width := flag.Int("width", defaultWidth, "The width of the Game of Life grid (in cells).")
	height := flag.Int("height", defaultHeight, "The height of the Game of Life grid (in cells).")
	chance := flag.Float64("chance", defaultChance, "The probability (0.0 to 1.0) that a cell starts alive.")
	tps := flag.Int("tps", defaultTPS, "The number of simulation updates per second (Ticks Per Second).")
	ruleStr := flag.String("rule", defaultRule, "Life-like cellular automaton rule in B.../S... format (e.g. B3/S23, B36/S23, B35678/S5678).")
	aliveHex := flag.String("alive", defaultAliveHex, "Hexadecimal RGBA color for alive cells.")
	deadHex := flag.String("dead", defaultDeadHex, "Hexadecimal RGBA color for dead cells.")
	flag.Parse()

	rule, err := NewRule(*ruleStr)
	if err != nil {
		return Config{}, err
	}
	alive, err := NewColor(*aliveHex)
	if err != nil {
		return Config{}, err
	}
	dead, err := NewColor(*deadHex)
	if err != nil {
		return Config{}, err
	}

	return NewConfig(*width, *height, *chance, *tps, rule, alive, dead)
}
