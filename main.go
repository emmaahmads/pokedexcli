package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/emmaahmads/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
	pokedex         pokeapi.Pokedex
}

func main() {
	client := pokeapi.NewClient(5 * time.Minute)
	pokedex := pokeapi.NewPokedex()

	cfg := &config{
		pokeapiClient: client,
		pokedex:       pokedex,
	}

	startRepl(cfg)
}

func startRepl(cfg *config) {
	fmt.Println("Pokedex >")
	option := ""
	scan := bufio.NewScanner(os.Stdin)
	for {
		for scan.Scan() {
			in := strings.ToLower(scan.Text())
			out := strings.Fields(in)
			cmdname, exists := getCommand()[out[0]]
			if len(out) > 1 {
				option = out[1]
			}

			if exists {
				err := cmdname.callback(cfg, option)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
