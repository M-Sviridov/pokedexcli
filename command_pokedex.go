package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex.Entries) == 0 {
		fmt.Println("you haven't caught any pokemons!")
	} else {
		fmt.Println("Your Pokedex:")
	}

	for _, pokemon := range cfg.pokedex.Entries {
		fmt.Printf("- %v\n", pokemon.Name)
	}
	return nil
}
