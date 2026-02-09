package main

import (
	"errors"
	"fmt"
)


func commandMap(cfg *config) error {
	locs, err := cfg.pokeapiClient.GetLocations(cfg.nextMapURL)
	if err != nil {
	    return err
	} 

	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	} 
	cfg.prevMapURL, cfg.nextMapURL = locs.Previous, locs.Next
	return nil
}


func commandMapb(cfg *config) error {
	if cfg.prevMapURL ==  nil {
		return errors.New("you're on the first page")
	} 
	locs, err := cfg.pokeapiClient.GetLocations(cfg.prevMapURL)
	if err != nil {
	    return err
	} 

	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	} 
	cfg.prevMapURL, cfg.nextMapURL = locs.Previous, locs.Next
	return nil
} 

