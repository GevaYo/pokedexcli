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
			name:        "explore",
			description: "Display the previous list of location from the Pokemon world",
			callback:    commandExplore,
			args:        []string{},
		},
		"catch": {
			name:        "catch",
			description: "Catching Pokemon adds them to the user's Pokedex",
			callback:    commandCatch,
			args:        []string{},
		},
		"inspect": {
			name:        "inspect",
			description: "Show details for pokemons in the Pokedex",
			callback:    commandInspect,
			args:        []string{},
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show all of the captured pokemons",
			callback:    commandPokedex,
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
	locationName := args[0]
	pokemons, err := config.pokeapiClient.ListPokemonsInLocationArea(&locationName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", locationName)
	for _, pokemon := range pokemons.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}

func commandCatch(config *Config, args []string) error {
	pokemoneName := args[0]
	pokemonData, err := config.pokeapiClient.FetchPokemonByName(&pokemoneName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemoneName)
	isCaught := config.Pokedex.AttemptCatch(pokemonData)
	if isCaught {
		config.Pokedex.AddToPokedex(pokemoneName, pokemonData)
		fmt.Printf("%s was caught!\n", pokemoneName)
	} else {
		fmt.Printf("%s escaped!\n", pokemoneName)
	}
	return nil
}

func commandInspect(config *Config, args []string) error {
	pokemonName := args[0]
	pokemonData, exists := config.Pokedex.caught[pokemonName]
	if !exists {
		fmt.Println("You have not caught that pokemon")
	} else {
		fmt.Printf("Name: %s\n", pokemonName)
		fmt.Printf("Height: %d\n", pokemonData.Height)
		fmt.Printf("Weight: %d\n", pokemonData.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemonData.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemonData.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}
	}
	return nil
}

func commandPokedex(config *Config, args []string) error {
	pokedex := config.Pokedex.caught
	if len(pokedex) == 0 {
		return fmt.Errorf("No pokemon were caught")
	}
	fmt.Println("Your Pokedex:")
	for key := range pokedex {
		fmt.Printf("  - %s\n", key)
	}
	return nil
}
