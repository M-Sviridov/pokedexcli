package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide one location name")
	}
	location := args[0]

	url := "https://pokeapi.co/api/v2/location-area/" + location
	encounters, err := cfg.pokeapiClient.LocationPokemonList(url)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")

	for _, encounter := range encounters.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
