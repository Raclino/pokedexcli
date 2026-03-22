package cli

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Raclino/pokedexcli/internal/pokeapi"
	"github.com/Raclino/pokedexcli/internal/pokecache"
)

// TODO: voir pour refacto via une appConfig global dans repl.go et update les signatrues des callbacks func
var client = &http.Client{Timeout: 3 * time.Second}
var cache = pokecache.NewCache(5 * time.Second)

func updateConfigFromResponse(config *pokeapi.LocationAreaConfig, resp *pokeapi.LocationAreasResponse) {
	if resp.Next != nil {
		config.Next = *resp.Next
	} else {
		config.Next = ""
	}

	if resp.Previous != nil {
		config.Previous = *resp.Previous
	} else {
		config.Previous = ""
	}
}

func commandMap(config *pokeapi.LocationAreaConfig, args ...string) error {
	resp, err := pokeapi.GetLocationAreas(client, cache, config.Next)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	updateConfigFromResponse(config, resp)
	return nil
}

func commandMapb(config *pokeapi.LocationAreaConfig, args ...string) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	resp, err := pokeapi.GetLocationAreas(client, cache, config.Previous)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	updateConfigFromResponse(config, resp)
	return nil
}
