package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Raclino/pokedexcli/internal/pokeapi"
)

const limit = 20

var client = &http.Client{Timeout: 3 * time.Second}
var startRange = 1

func commandMap(config *Config) error {
	fmt.Println(*config)
	endRange := startRange + limit - 1

	err := pokeapi.FetchLocationsNames(client, startRange, endRange)
	if err != nil {
		return err
	}

	startRange += limit
	return nil
}

func commandMapb(config *Config) error {
	fmt.Println(*config)
	startRange -= limit * 2
	endRange := startRange + limit - 1

	err := pokeapi.FetchLocationsNames(client, startRange, endRange)
	if err != nil {
		return err
	}

	return nil
}
