package main

import (
	"fmt"
	"os"

	"github.com/Raclino/pokedexcli/internal/pokeapi"
)

func commandExit(config *pokeapi.LocationAreaConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
