package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
	args        []string
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Describe how to use the Pokedex",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display a list of location from the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous list of location from the Pokemon world",
			callback:    commandMapB,
		},
		"explore": {
			name:        "mapb",
			description: "Display the previous list of location from the Pokemon world",
			callback:    commandExplore,
			args:        []string{},
		},
	}

}

func commandExit(config *Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	fmt.Println()
	os.Exit(0)
	return nil
}

func commandHelp(config *Config, args []string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	for cmdName := range commands {
		command := commands[cmdName]
		fmt.Println(fmt.Sprintf("%s : %s", command.name, command.description))
	}

	fmt.Println()
	return nil
}

func commandMap(config *Config, args []string) error {
	locationResp, err := config.pokeapiClient.ListLocationAreas(config.nextURL)
	if err != nil {
		return err
	}

	config.nextURL = locationResp.Next
	config.prevURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapB(config *Config, args []string) error {
	if config.prevURL == nil {
		return errors.New("Already on the first page!")
	}
	locationResp, err := config.pokeapiClient.ListLocationAreas(config.prevURL)
	if err != nil {
		return err
	}

	config.nextURL = locationResp.Next
	config.prevURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExplore(config *Config, args []string) error {
	return nil
}
