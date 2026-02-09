package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, parameters ...string) error {
	if len(parameters) != 1 {
		return errors.New("No location passed. Use `map <location>`")
	} 
	loc, err := cfg.pokeapiClient.ExploreLocation(parameters[0])
	if err != nil {
	    return err
	} 

	fmt.Printf("Exploring %s...\n", loc.Name)
	fmt.Println("Found Pokemon: ")
	for _, p := range loc.PokemonEncounters {
		fmt.Println(" - ", p.Pokemon.Name)
	} 
	return nil
} 
