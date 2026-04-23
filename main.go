package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanText := cleanInput(input)
		value, ok := commands[cleanText[0]]
		if ok {
			err := value.callback()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Print("Unknown command\n")
		}
	}
}
