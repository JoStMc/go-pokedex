package main

import (
	"errors"
	"fmt"
)


func commandMap(cfg *config, parameters ...string) error {
	locs, err := cfg.pokeapiClient.ListLocations(cfg.nextMapURL)
	if err != nil {
	    return err
	} 

	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	} 
	cfg.prevMapURL, cfg.nextMapURL = locs.Previous, locs.Next
	return nil
}


func commandMapb(cfg *config, parameters ...string) error {
	if cfg.prevMapURL ==  nil {
		return errors.New("you're on the first page")
	} 
	locs, err := cfg.pokeapiClient.ListLocations(cfg.prevMapURL)
	if err != nil {
	    return err
	} 

	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	} 
	cfg.prevMapURL, cfg.nextMapURL = locs.Previous, locs.Next
	return nil
} 

