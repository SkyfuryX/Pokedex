package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n")
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
		return errors.New("you're on the first page")
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
