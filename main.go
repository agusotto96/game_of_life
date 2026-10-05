package main

import (
	"log"
)

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	world := NewWorld(
		config.Width,
		config.Height,
		config.Chance,
		config.Rule,
	)
	game := NewGame(
		world,
		config.TPS,
		config.Alive,
		config.Dead,
	)
	err = RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
