package main

import "fmt"

func commandInspect(cfg *config, option string) error {

	if ok := cfg.pokedex.GetPokemon(option); !ok {
		return fmt.Errorf("you have not caught %s", option)
	} else {
		cfg.pokedex.InspectPokemon(option)
	}
	return nil
}

func commandPokedex(cfg *config, option string) error {

	cfg.pokedex.ListPokedex()
	return nil
}
