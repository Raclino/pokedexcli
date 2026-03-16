package main

import (
	"fmt"
	"slices"
)

func commandHelp(config *Config) error {
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
