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
	callback func(*config, ...string) error
} 

type config struct {
	pokeapiClient pokeapi.Client
    nextMapURL *string
    prevMapURL *string
	caughtPokemon map[string]pokeapi.Pokemon
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
		parameters := []string{}
		if len(input) > 1 {
			parameters = input[1:]
		} 
		command, exists := getCommands()[input[0]]
		if exists {
			err := command.callback(cfg, parameters...)
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
		"explore": {
			name: "explore <location_name>",
			description: "List the Pokemon found in a location from `map`",
			callback: commandExplore,
		},
		"catch": {
			name: "catch <pokemon_name>",
			description: "Throw a Pokeball at a Pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect <pokemon_name>",
			description: "Check the stats of a caught Pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "List which Pokemon you've caught",
			callback: commandPokedex,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	} 
} 
