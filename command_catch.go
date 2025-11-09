package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide one pokemon name")
	}
	pokemon := args[0]

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	pokemonStats, err := cfg.pokeapiClient.GetPokemonStats(url)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	success := attemptCatch(pokemonStats.BaseExperience)

	if success {
		fmt.Printf("%s was caught!\n", pokemon)
		cfg.pokedex.Entries[pokemon] = pokemonStats
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemon)

	return nil
}

func attemptCatch(baseExperience int) bool {
	catchRate := 100 - (baseExperience / 3)
	if catchRate < 25 {
		catchRate = 25
	}

	roll := rand.Intn(100)
	return roll < catchRate
}
