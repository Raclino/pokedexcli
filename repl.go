package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Raclino/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config) error
}

func startRepl() {
	urlsConfig := &pokeapi.Config{Previous: pokeapi.LocationsAreas + strconv.Itoa(1)}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		newInputText := scanner.Text()
		cleanedInput := cleanInput(newInputText)

		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]

		c, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := c.callback(urlsConfig)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	loweredText := strings.ToLower(trimmedText)

	inputCleaned := []string{}

	for w := range strings.FieldsSeq(loweredText) {
		inputCleaned = append(inputCleaned, w)
	}

	return inputCleaned
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
			description: "Displays the names of 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous names of 20 location areas in the Pokemon world.",
			callback:    commandMapb,
		},
	}
}
