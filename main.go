package main

import (
	"github.com/M-Sviridov/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second*5, time.Minute*5)
	cfg := config{
		pokedex: pokeapi.Pokedex{
			Entries: make(map[string]pokeapi.RespPokemon),
		},
		pokeapiClient: pokeClient,
	}
	startRepl(&cfg)
}
