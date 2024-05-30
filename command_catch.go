package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, option string) error {
	name := option
	pokemon, err := cfg.pokeapiClient.GetPokemon(&name)
	if err != nil {
		return err
	}
	chance := rand.Intn(pokemon.BaseExperience)

	if chance > pokemon.BaseExperience/2 {
		fmt.Println(" yay, you have caught " + name)
		cfg.pokedex.AddPokemon(pokemon)
	} else {
		fmt.Println(" oops, you did not catch " + name)
	}
	return nil
}
