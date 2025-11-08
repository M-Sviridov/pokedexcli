package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, area string) error {
	locations, err := cfg.pokeapiClient.LocationAreaList(cfg.Next)
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

func commandMapb(cfg *config, area string) error {
	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	}

	locations, err := cfg.pokeapiClient.LocationAreaList(cfg.Previous)
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
