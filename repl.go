package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/M-Sviridov/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	pokedex       pokeapi.Pokedex
	Next          *string
	Previous      *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Explore <area_name> and list pokemons",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect",
			description: "Display stats of a caught Pokemon",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of 20 previous location areas",
			callback:    commandMapb,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		input := cleanInput(scanner.Text())
		args := []string{}
		if len(input) == 0 {
			continue
		}

		if len(input) > 1 {
			args = input[1:]
		}

		commandName := input[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
