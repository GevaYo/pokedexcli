package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowered := strings.ToLower(trimmed)
	words := strings.Fields(lowered)

	return words
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print("Pokedex > ")
		text := scanner.Text()
		if text == "" {
			continue
		}
		inputs := cleanInput(text)
		fmt.Println("Your command was: " + inputs[0])
	}
}
