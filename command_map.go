package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Raclino/pokedexcli/internal/pokeapi"
)

var client = &http.Client{Timeout: 3 * time.Second}

func commandMap(config *pokeapi.LocationAreaConfig) error {
	resp, err := pokeapi.GetLocationAreas(client, config.Next)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

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

	return nil
}

func commandMapb(config *pokeapi.LocationAreaConfig) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	resp, err := pokeapi.GetLocationAreas(client, config.Previous)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

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

	return nil
}
