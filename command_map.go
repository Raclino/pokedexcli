package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type locationAreaAPIResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

const locationsAreas = "https://pokeapi.co/api/v2/location-area"
const limit = 20

var client = &http.Client{Timeout: 3 * time.Second}
var startRange = 1

func commandMap() error {
	endRange := startRange + limit - 1
	fetchLocationsNames(client, startRange, endRange)
	startRange += limit
	return nil
}

func commandMapb() error {
	startRange -= limit * 2
	endRange := startRange + limit - 1
	fetchLocationsNames(client, startRange, endRange)
	return nil
}

func fetchLocationsNames(client *http.Client, startRange, endRange int) error {

	for i := startRange; i <= endRange; i++ {
		fullURL := locationsAreas + "/" + strconv.Itoa(i)

		req, err := http.NewRequest(http.MethodGet, fullURL, nil)
		if err != nil {
			return fmt.Errorf("failed to create request for %s: %w", fullURL, err)
		}

		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("failed to fetch page %d: %w", i+1, err)
		}

		location := locationAreaAPIResponse{}
		if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
			return fmt.Errorf("Couldn't decode response body: %w", err)
		}

		resp.Body.Close()

		fmt.Println(location.Location.Name)
	}
	return nil
}
