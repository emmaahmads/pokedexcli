package main

import (
	"fmt"
	"log"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func commandMap(cfg *config) error {
	// get location list

	locations, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err

	}

	cfg.nextLocationURL = locations.Next
	cfg.prevLocationURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	log.Println("test commandMapb")
	if cfg.prevLocationURL == nil {
		log.Println("test commandMapb 1")
		return fmt.Errorf("no previous location")
	}

	// get location list
	log.Println("test commandMapb 2")
	locations, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	log.Println("test commandMapb 3")
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locations.Next
	cfg.prevLocationURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
