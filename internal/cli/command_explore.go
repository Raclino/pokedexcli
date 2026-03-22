package cli

import (
	"fmt"

	"github.com/Raclino/pokedexcli/internal/pokeapi"
)

func commandExplore(config *pokeapi.LocationAreaConfig, args ...string) error {
	locationArea := args[1]
	fmt.Printf("Exploring %s...\n", locationArea)

	resp, err := pokeapi.GetLocationAreaInfos(client, cache, locationArea)
	if err != nil {
		return err
	}

	for _, pokeEncounter := range resp.PokemonEncounters {
		fmt.Println(pokeEncounter.Pokemon.Name)
	}

	return nil
}
