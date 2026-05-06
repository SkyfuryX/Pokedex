package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/SkyfuryX/pokedex/internal"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	split := strings.Fields(lower)
	return split
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	input []string
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex map[string]pokeapi.Pokemon
}

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		cfg.input = cleanInput(input)
		value, ok := commands[cfg.input[0]]
		if ok {
			err := value.callback(cfg)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		} else {
			fmt.Print("Unknown command\n")
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Shows the next 20 locations from the Pokemon World",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous 20 locations from the Pokemon World",
			callback:    commandMapb,
		},
		"explore": {
			name: "explore <area-name>",
			description: "Shows pokemon that can be found within the area",
			callback: explore,
		},
		"catch": {
			name: "catch",
			description: "Attempt to catch a pokemon",
			callback: catch,
		},
		"inspect": {
			name: "inspect",
			description: "View a pokemon you've caught",
			callback: inspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "View what pokemon you've caught",
			callback: pokedex,
		},
	}
}
