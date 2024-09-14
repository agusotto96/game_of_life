package main

import (
	"log"
)

func main() {
	config := ReadConfig()
	world := RandomWorld(
		config.Width,
		config.Height,
		config.Chance,
	)
	game := NewGame(world)
	err := game.Run()
	if err != nil {
		log.Fatal(err)
	}
}
