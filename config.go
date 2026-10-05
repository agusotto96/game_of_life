package main

import (
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
)

const (
	defaultWidth    = 960
	defaultHeight   = 600
	defaultChance   = 0.15
	defaultTPS      = 60
	defaultRule     = "B3/S23"
	defaultAliveHex = "39ff14ff"
	defaultDeadHex  = "202020ff"
	hexColorLength  = 8
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

func ReadConfig() (Config, error) {
	width := flag.Int("width", defaultWidth, "The width of the Game of Life grid (in cells).")
	height := flag.Int("height", defaultHeight, "The height of the Game of Life grid (in cells).")
	chance := flag.Float64("chance", defaultChance, "The probability (0.0 to 1.0) that a cell starts alive.")
	tps := flag.Int("tps", defaultTPS, "The number of simulation updates per second (Ticks Per Second).")
	ruleStr := flag.String("rule", defaultRule, "Life-like cellular automaton rule in B.../S... format (e.g. B3/S23, B36/S23, B35678/S5678).")
	aliveHex := flag.String("alive", defaultAliveHex, "Hexadecimal RGBA color for alive cells.")
	deadHex := flag.String("dead", defaultDeadHex, "Hexadecimal RGBA color for dead cells.")
	flag.Parse()

	rule, err := ParseRule(*ruleStr)
	if err != nil {
		return Config{}, err
	}
	alive, err := parseColor(*aliveHex)
	if err != nil {
		return Config{}, err
	}
	dead, err := parseColor(*deadHex)
	if err != nil {
		return Config{}, err
	}
	config := Config{
		Width:  *width,
		Height: *height,
		Chance: *chance,
		TPS:    *tps,
		Rule:   rule,
		Alive:  alive,
		Dead:   dead,
	}
	if err := validateConfig(config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func validateConfig(c Config) error {
	if c.Width <= 0 {
		return errors.New("width must be greater than 0")
	}
	if c.Height <= 0 {
		return errors.New("height must be greater than 0")
	}
	if c.Chance <= 0 || c.Chance > 1 {
		return errors.New("chance must be between 0.0 and 1.0")
	}
	if c.TPS <= 0 {
		return errors.New("tps must be greater than 0")
	}
	return nil
}

func parseColor(hexStr string) (Color, error) {
	if len(hexStr) != hexColorLength {
		return Color{}, fmt.Errorf("invalid hexadecimal color format, expected %d characters (RRGGBBAA)", hexColorLength)
	}
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return Color{}, err
	}
	return (Color)(bytes), nil
}
