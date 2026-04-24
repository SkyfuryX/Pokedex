package main

import (
	pokeapi "github.com/SkyfuryX/pokedex/internal"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startREPL(cfg)
}
