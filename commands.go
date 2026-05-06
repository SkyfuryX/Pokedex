package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

func commandMapf(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = &locationsResp.Next
	cfg.prevLocationsURL = &locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("You're already on the first page")
	}
	locationsResp, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = &locationsResp.Next
	cfg.prevLocationsURL = &locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}

func explore(cfg *config) error {
	if len(cfg.input) < 2 {
		return errors.New("Must include a location, ex: explore canalave-city-area")
	}
	PageURL := "https://pokeapi.co/api/v2/location-area/" + cfg.input[1]
	areaResp, err := cfg.pokeapiClient.GetArea(&PageURL)

	if err != nil {
		return errors.New("Invalid location")
	}

	fmt.Printf("Exploring %v...\n", cfg.input[1])
	for _, encounter := range areaResp.Encounters {
		fmt.Printf("- %v\n", encounter.Pokemon.Name)
	}
	return nil
}

func catch(cfg *config) error {
	if len(cfg.input) < 2 {
		return errors.New("Must include the name of the Pokemon you want to catch")
	}
	pageURL := "https://pokeapi.co/api/v2/pokemon/" + cfg.input[1]
	pokemon, err := cfg.pokeapiClient.GetPokemon(&pageURL)
	if err != nil {
		return errors.New("Pokemon not found")
	}
	rand := rand.Intn(101) // number between 0-100
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)

	exp := pokemon.BaseExperience //handles pokemon that have very high base exp (looking at you, Blissey)
	if exp >= 360 {
		exp = 360
	}
	if rand > (exp / 4) {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = pokemon //adds to pokedex once caught
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}
	return nil
}

func inspect(cfg *config) error {
	if len(cfg.input) < 2 {
		return errors.New("Choose a pokemon to view. Use 'pokedex' to see which pokemon you've caught.")
	}
	pokemon, ok := cfg.pokedex[cfg.input[1]]
	if !ok {
		return errors.New("You haven't caught this pokemon yet")
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Print("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Print("Types:\n")
	for _, ty := range pokemon.Types {
		fmt.Printf("  - %v\n", ty.Type.Name)
	}
	return nil
}

func pokedex(cfg *config) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("You haven't caught any pokemon yet!")
	}
	fmt.Print("Your Pokedex:\n")
	for key := range cfg.pokedex {
		fmt.Printf(" - %v\n", key)
	}
	return nil
}
