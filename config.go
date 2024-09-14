package main

import "flag"

type Config struct {
	Width  int
	Height int
	Chance int
}

func ReadConfig() Config {
	width := flag.Int("width", 640, "The width of the Game of Life grid (in cells).")
	height := flag.Int("height", 480, "The height of the Game of Life grid (in cells).")
	chance := flag.Int("chance", 15, "The probability (1 in X) that a cell starts alive. Lower values increase the number of alive cells.")
	flag.Parse()
	return Config{
		Width:  *width,
		Height: *height,
		Chance: *chance,
	}
}
