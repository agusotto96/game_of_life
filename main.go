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
	worlds := UpdateWorlds(world)
	game := NewGame(
		worlds,
		config.Width,
		config.Height,
		config.Alive,
		config.Dead,
	)
	err = RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
