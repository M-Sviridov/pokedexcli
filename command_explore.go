package main

import "fmt"

func commandExplore(cfg *config, area string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + area
	encounters, err := cfg.pokeapiClient.LocationPokemonList(url)
	if err != nil {
		return err
	}

	for _, encounter := range encounters.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
