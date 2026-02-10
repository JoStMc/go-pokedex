package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, parameters ...string) error {
	if len(parameters) != 1 {
		return errors.New("no name passed. use `catch <pokemon>`")
	} 
	pok, err := cfg.pokeapiClient.GetPokemon(parameters[0])
	if err != nil {
	    return err
	} 

	fmt.Printf("Throwing a Pokeball at %s...\n", pok.Name)

	if rand.Intn(pok.BaseExperience) < 36 {
		fmt.Printf("%s was caught!\n", pok.Name)
		cfg.caughtPokemon[pok.Name] = pok
	} else {
	    fmt.Printf("%s escaped!\n", pok.Name)
	} 

	return nil
} 
