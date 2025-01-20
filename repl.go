package main

import (
	"bufio"
	"fmt"
	"os"
	pokeapi "pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokecache"
	"strings"
)

type Config struct {
	pokeapiClient *pokeapi.Client
	pokeCache     *pokecache.Cache
	Pokedex       *Pokedex
	nextURL       *string
	prevURL       *string
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowered := strings.ToLower(trimmed)
	words := strings.Fields(lowered)

	return words
}

func startRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		input := getInput(scanner)
		cleanInput := cleanInput(input)
		if len(cleanInput) == 0 {
			continue
		}
		command := cleanInput[0]
		args := []string{}
		if len(cleanInput) > 1 {
			args = parseArgs(cleanInput)
		}
		if cmd, exists := commands[command]; exists {
			if err := validateCommand(command, args); err != nil {
				fmt.Println(err)
				continue
			}
			err := cmd.callback(config, args)
			if err != nil {
				fmt.Println(fmt.Sprintf("Error executing command: %s ", cmd.name))
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func validateCommand(command string, args []string) error {
	switch command {
	case "help", "exit", "map", "mapb", "pokedex":
		if len(args) > 0 {
			return fmt.Errorf("%v command doesn't accept arguments", command)
		}
	case "explore", "catch", "inspect":
		if len(args) != 1 {
			return fmt.Errorf("%v command accept exactly 1 argument", command)
		}
	}
	return nil
}
func getInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func parseArgs(inputs []string) []string {
	return inputs[1:]
}
