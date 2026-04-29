package main

import (
	"time"

	pokeapi "github.com/SkyfuryX/pokedex/internal"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second, 30 * time.Second),
	}
	startREPL(cfg)
}
