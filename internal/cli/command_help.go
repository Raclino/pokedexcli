package cli

import (
	"fmt"
	"slices"

	"github.com/Raclino/pokedexcli/internal/pokeapi"
)

func commandHelp(config *pokeapi.LocationAreaConfig, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")

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
