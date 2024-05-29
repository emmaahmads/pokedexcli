package main

import (
	"fmt"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, option string) error
}

func commandMap(cfg *config, option string) error {
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

func commandMapb(cfg *config, option string) error {
	if cfg.prevLocationURL == nil {
		return fmt.Errorf("no previous location")
	}

	// get location list
	locations, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
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
