package main

import (
	"errors"
	"fmt"

	"github.com/M-Sviridov/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide one pokemon name")
	}
	pokemon := args[0]

	if stats, ok := cfg.pokedex.Entries[pokemon]; ok {
		displayStats(stats)
	} else {
		fmt.Println("you did not caught this pokemon")
	}
	return nil
}

func displayStats(stats pokeapi.RespPokemon) {
	fmt.Printf("Name: %s\n", stats.Name)
	fmt.Printf("Height: %d\n", stats.Height)
	fmt.Printf("Weight: %d\n", stats.Weight)
	fmt.Println("Stats:")
	for _, stat := range stats.Stats {
		fmt.Printf("- %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range stats.Types {
		fmt.Printf("- %s\n", pokemonType.Type.Name)
	}
}
