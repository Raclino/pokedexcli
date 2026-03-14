package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	commandNames := make([]string, 0, len(getCommands()))
	for name := range getCommands() {
		commandNames = append(commandNames, name)
	}
	slices.Sort(commandNames)

	for _, name := range commandNames {
		command := getCommands()[name]
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
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
	}
}
