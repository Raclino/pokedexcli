package cli

import (
	"fmt"
	"math/rand"

	"github.com/Raclino/pokedexcli/internal/pokeapi"
)

func commandCatch(appConfig *AppConfig, args ...string) error {

	pokemonName := args[1]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	resp, err := pokeapi.GetPokemonInfos(appConfig.client, pokemonName)
	if err != nil {
		return err
	}

	isCaught := tryCatchPokemon(resp.BaseExperience)
	if !isCaught {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	appConfig.pokedex[pokemonName] = *resp
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

	randomNumber := rand.Intn(95)

	if randomNumber > catchChance {
		return false

	}

	return true
}
