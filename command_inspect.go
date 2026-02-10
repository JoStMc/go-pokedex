package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, parameters ...string) error {
	if len(parameters) != 1 {
		return errors.New("no name passed. use `inspect <pokemon>`")
	} 
	pok, caught := cfg.caughtPokemon[parameters[0]]
	if !caught {
	    fmt.Println("you have not caught that pokemon")
		return nil
	} 

	fmt.Println("Name:", pok.Name)
	fmt.Println("Height:", pok.Height)
	fmt.Println("Weight:", pok.Weight)
	fmt.Println("Stats:")
	for _, stat := range pok.Stats {
		fmt.Printf("  - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pok.Types {
		fmt.Println("  - ", typ.Type.Name)
	} 

	return nil
}
