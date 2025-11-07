package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type LocationAreaList struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*LocationAreaList) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
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

func startRepl() {
	locations, err := initialiseLocationAreaList()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(locations)
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

func commandExit(locations *LocationAreaList) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(locations *LocationAreaList) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func initialiseLocationAreaList() (*LocationAreaList, error) {
	url := "https://pokeapi.co/api/v2/location-area/"

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var locationAreas *LocationAreaList
	if err := json.Unmarshal(body, &locationAreas); err != nil {
		return nil, err
	}

	return locationAreas, nil
}

func updateLocationAreaList(command string, locations *LocationAreaList) error {
	url := ""
	switch command {
	case "map":
		url = *locations.Next
	case "mapb":
		url = *locations.Previous
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := json.Unmarshal(body, &locations); err != nil {
		return err
	}

	return nil
}

func commandMap(locations *LocationAreaList) error {
	if locations.Next == nil {
		fmt.Println("you're on the last page")
	} else {
		for _, location := range locations.Results {
			fmt.Println(location.Name)
		}
		updateLocationAreaList("map", locations)
	}
	return nil
}

func commandMapb(locations *LocationAreaList) error {
	if locations.Previous == nil {
		fmt.Println("you're on the first page")
	} else {
		for _, location := range locations.Results {
			fmt.Println(location.Name)
		}
		updateLocationAreaList("mapb", locations)
	}
	return nil
}
