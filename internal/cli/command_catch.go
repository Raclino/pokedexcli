package cli

import (
	"fmt"
	"math/rand"

	"github.com/Raclino/pokedexcli/internal/pokeapi"
)

var Pokedex = map[string]pokeapi.PokemonInfos{}

func commandCatch(config *pokeapi.LocationAreaConfig, args ...string) error {

	pokemonName := args[1]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	resp, err := pokeapi.GetPokemonInfos(client, pokemonName)
	if err != nil {
		return err
	}
	fmt.Println("baseExperience: ", resp.BaseExperience)

	isCaught := tryCatchPokemon(resp.BaseExperience)
	fmt.Println(isCaught)
	if !isCaught {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	Pokedex[pokemonName] = *resp
	fmt.Printf("%s was caught!\n", pokemonName)
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
