package cli

import (
	"fmt"
	"os"

	"github.com/Raclino/pokedexcli/internal/pokeapi"
)

func commandExit(config *pokeapi.LocationAreaConfig, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
