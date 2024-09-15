package main

import (
	"log"
)

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	world := RandomWorld(
		config.Width,
		config.Height,
		config.Chance,
	)
	game := NewGame(
		world,
		config.Alive,
		config.Dead,
	)
	err = RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
