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
	Alive  Color
	Dead   Color
}

func ReadConfig() (Config, error) {
	width := flag.Int("width", 640, "The width of the Game of Life grid (in cells).")
	height := flag.Int("height", 480, "The height of the Game of Life grid (in cells).")
	chance := flag.Int("chance", 15, "The probability (1 in X) that a cell starts alive. Lower values increase the number of alive cells.")
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
		Alive:  alive,
		Dead:   dead,
	}
	return config, nil
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
