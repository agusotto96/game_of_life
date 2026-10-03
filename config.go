package main

import (
	"encoding/hex"
	"errors"
	"flag"
)

type Config struct {
	Width  int
	Height int
	Chance int
	TPS    int
	Alive  Color
	Dead   Color
}

func ReadConfig() (Config, error) {
	width := flag.Int("width", 640, "The width of the Game of Life grid (in cells).")
	height := flag.Int("height", 480, "The height of the Game of Life grid (in cells).")
	chance := flag.Int("chance", 15, "The probability (1 in X) that a cell starts alive. Lower values increase the number of alive cells.")
	tps := flag.Int("tps", 60, "The number of simulation updates per second (Ticks Per Second).")
	aliveHex := flag.String("alive", "39ff14ff", "Hexadecimal RGBA color for alive cells.")
	deadHex := flag.String("dead", "202020ff", "Hexadecimal RGBA color for dead cells.")
	flag.Parse()
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
	if c.Chance <= 0 {
		return errors.New("chance must be greater than 0")
	}
	if c.TPS <= 0 {
		return errors.New("tps must be greater than 0")
	}
	return nil
}

func parseColor(hexStr string) (Color, error) {
	if len(hexStr) != 8 {
		return Color{}, errors.New("invalid hexadecimal color format, expected 8 characters (RRGGBBAA)")
	}
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return Color{}, err
	}
	return (Color)(bytes), nil
}
