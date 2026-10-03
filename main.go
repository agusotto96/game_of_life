package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	world := RandomWorld(
		config.Width,
		config.Height,
		config.Chance,
	)
	worlds := UpdateWorlds(ctx, world)
	game := NewGame(
		worlds,
		config.Width,
		config.Height,
		config.TPS,
		config.Alive,
		config.Dead,
	)
	err = RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
