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
	fmt.Print("Pokedex > ")
	for {
		input := getInput()
		if input == "" {
			continue
		}
		cleanInput := cleanInput(input)
		command, args := parseInput(cleanInput)
		if cmd, exists := commands[command]; exists {
			err := cmd.callback(config, args)
			if err != nil {
				fmt.Println(fmt.Sprintf("Error executing command: %s ", cmd.name))
			}
		} else {
			fmt.Println("Unknown command")
		}
		fmt.Print("Pokedex > ")
	}
}

func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func parseInput(inputs []string) (string, []string) {
	return inputs[0], inputs[1:]
}
