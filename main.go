package main

import (
	"log"
)

func main() {
	width := 640
	height := 480
	chance := 15
	world := NewWorld(width, height, chance)
	game := NewGame(world)
	err := game.Run()
	if err != nil {
		log.Fatal(err)
	}
}
