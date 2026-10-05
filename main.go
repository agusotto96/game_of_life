package main

import (
	"log"
)

func main() {
	c, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	w := NewWorld(c.Width, c.Height, c.Chance, c.Rule)
	g := NewGame(w, c.TPS, c.Alive, c.Dead)
	err = g.Run()
	if err != nil {
		log.Fatal(err)
	}
}
