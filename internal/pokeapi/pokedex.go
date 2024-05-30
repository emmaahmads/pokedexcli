package pokeapi

import "fmt"

type Pokedex struct {
	pokemon map[string]RespPokemon
}

func NewPokedex() Pokedex {
	return Pokedex{pokemon: make(map[string]RespPokemon)}
}

func (p *Pokedex) AddPokemon(pokemon RespPokemon) {
	p.pokemon[pokemon.Name] = pokemon
}

func (p *Pokedex) GetPokemon(pokemon string) bool {
	if _, ok := p.pokemon[pokemon]; ok {
		return true
	}
	return false
}

func (p *Pokedex) InspectPokemon(pokemon string) {
	fmt.Println("Name:", p.pokemon[pokemon].Name)
	fmt.Println("Height:", p.pokemon[pokemon].Height)
	fmt.Println("Weight:", p.pokemon[pokemon].Weight)
	fmt.Println("Stats:")
	for _, stat := range p.pokemon[pokemon].Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range p.pokemon[pokemon].Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}
}

func (p *Pokedex) ListPokedex() {
	for _, list := range p.pokemon {
		p.InspectPokemon(list.Name)
	}
}
