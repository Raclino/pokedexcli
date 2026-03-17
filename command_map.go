package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Raclino/pokedexcli/internal/pokeapi"
)

const limit = 20

var client = &http.Client{Timeout: 3 * time.Second}
var startRange = 1

func commandMap(config *pokeapi.Config) error {
	endRange := startRange + limit - 1
	config.Next = pokeapi.LocationsAreas + strconv.Itoa(endRange)
	fmt.Printf("PreviousLink: %s, NextLink: %s \n", config.Previous, config.Next)

	err := pokeapi.FetchLocationsNames(client, config, startRange, endRange)
	if err != nil {
		return err
	}

	startRange += limit
	config.Previous = pokeapi.LocationsAreas + strconv.Itoa(startRange)

	return nil
}

func commandMapb(config *pokeapi.Config) error {
	fmt.Println(*config)
	startRange -= limit * 2
	endRange := startRange + limit - 1

	err := pokeapi.FetchLocationsNames(client, config, startRange, endRange)
	if err != nil {
		return err
	}

	return nil
}
