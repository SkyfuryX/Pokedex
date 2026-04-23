package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	split := strings.Fields(text)
	for i, word := range split {
		split[i] = strings.ToLower(word)
	}
	return split
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n")
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

