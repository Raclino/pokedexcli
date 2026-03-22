package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Raclino/pokedexcli/internal/pokecache"
)

type LocationAreaConfig struct {
	Next     string
	Previous string
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreasResponse struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type PokemonEncounters struct {
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
}
type LocationAreaInfosResponse struct {
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
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

const LocationsAreas string = "https://pokeapi.co/api/v2/location-area/"

func GetLocationAreas(client *http.Client, cache *pokecache.Cache, url string) (*LocationAreasResponse, error) {
	var body []byte
	var ok bool

	body, ok = cache.Get(url)
	if !ok {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request for %s: %w", url, err)
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch location areas: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}

		cache.Add(url, body)
	}

	locationAreasResponse := LocationAreasResponse{}
	if err := json.Unmarshal(body, &locationAreasResponse); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal response body: %w", err)
	}

	return &locationAreasResponse, nil
}

func GetLocationAreaInfos(client *http.Client, cache *pokecache.Cache, locationArea string) (*LocationAreaInfosResponse, error) {
	var body []byte
	var ok bool

	url := LocationsAreas + locationArea

	body, ok = cache.Get(url)
	if !ok {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request for %s: %w", url, err)
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch location area informations: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}

		cache.Add(url, body)
	}

	locationAreaInfosResponse := LocationAreaInfosResponse{}
	if err := json.Unmarshal(body, &locationAreaInfosResponse); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal response body: %w", err)
	}

	return &locationAreaInfosResponse, nil
}
