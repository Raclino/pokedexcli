package cli

import (
	"fmt"
	"math/rand"

	"github.com/Raclino/pokedexcli/internal/pokeapi"
)

// TODO: voir pour initialize le pokedex autre part / plus opti / logic
var Pokedex = map[string]pokeapi.PokemonInfos{}

func commandCatch(config *pokeapi.LocationAreaConfig, args ...string) error {

	pokemonName := args[1]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	resp, err := pokeapi.GetPokemonInfos(client, pokemonName)
	if err != nil {
		return err
	}

	isCaught := tryCatchPokemon(resp.BaseExperience)
	if !isCaught {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	Pokedex[pokemonName] = *resp
	fmt.Printf("%s was caught!\n", pokemonName)
	fmt.Println("You may now inspect it with the inspect command.")
	return nil
}

func tryCatchPokemon(baseExperience int) bool {
	catchChance := 100 - baseExperience

	if catchChance < 10 {
		catchChance = 10
	}

	if catchChance > 90 {
		catchChance = 90
	}

	randomNumber := rand.Intn(100)

	if randomNumber < catchChance {
		return true

	}

	return false
}
