package cli

import (
	"fmt"

	"github.com/Raclino/pokedexcli/internal/pokeapi"
)

func commandPokedex(config *pokeapi.LocationAreaConfig, args ...string) error {

	fmt.Println("Your Pokedex:")

	for _, p := range Pokedex {
		fmt.Printf("  - %s\n", p.Name)
	}
	return nil
}
