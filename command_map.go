package main

import (
	"errors"
	"fmt"
	"github.com/M-Sviridov/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	locations, err := pokeapi.LocationAreaList(cfg.Next)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	}

	locations, err := pokeapi.LocationAreaList(cfg.Previous)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	return nil
}
