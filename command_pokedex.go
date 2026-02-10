package main

import "fmt"

func commandPokedex(cfg *config, parameters ...string) error {
    if len(cfg.caughtPokemon) == 0 {
        fmt.Println("You have caught no pokemon.")
		return nil
    } 
	fmt.Println("Your Pokedex:")
	for _, pok := range cfg.caughtPokemon {
	    fmt.Println("  - ", pok.Name)
	} 
	return nil
} 
