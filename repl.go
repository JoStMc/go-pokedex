package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JoStMc/go-pokedex/internal/pokeapi"
)

type cliCommand struct {
    name string
	description string
	callback func(*config) error
} 

type config struct {
	pokeapiClient pokeapi.Client
    nextMapURL *string
    prevMapURL *string
} 

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
		    continue
		} 
		command, exists := getCommands()[input[0]]
		if exists {
			err := command.callback(cfg)
			if err != nil {
			    fmt.Println(err)
			} 
			continue
		} else {
		    fmt.Println("Unknown command")
			continue
		} 
	}
} 

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Returns the next 20 locations on the map",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Returns the previous 20 locations on the map",
			callback: commandMapb,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	} 
} 
