package main

import (
	"fmt"
)

func commandExplore(cfg *config, option string) error {
	explore, err := cfg.pokeapiClient.ExploreLocations(&option)
	if err != nil {
		return err
	}

	for _, list := range explore.PokemonEncounters {
		fmt.Println(string("\033[33m"), list.Pokemon.Name, string("\033[0m"))
	}

	return nil
}
