package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
    name string
	description string
	callback func() error
} 

func startRepl() {
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
			err := command.callback()
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
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	} 
} 
