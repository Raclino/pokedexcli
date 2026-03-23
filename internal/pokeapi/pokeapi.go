package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Raclino/pokedexcli/internal/pokecache"
)

const pokeAPIBaseURL string = "https://pokeapi.co/api/v2/"

const (
	LocationAreasURL = pokeAPIBaseURL + "location-area/"
	PokemonURL       = pokeAPIBaseURL + "pokemon/"
)

// TODO: voir pour refactorise la reretition au niveau de la construction et parsing des requetes externes

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

	url := LocationAreasURL + locationArea

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
		return nil, fmt.Errorf("couldn't Unmarshal response body: %w", err)
	}

	return &locationAreaInfosResponse, nil
}

func GetPokemonInfos(client *http.Client, pokemonName string) (*PokemonInfos, error) {

	url := PokemonURL + pokemonName

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for %s: %w", url, err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch %s informations: %w", pokemonName, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	PokemonInfos := PokemonInfos{}
	if err := json.NewDecoder(resp.Body).Decode(&PokemonInfos); err != nil {
		return nil, fmt.Errorf("couldn't decode response body: %w", err)
	}
	return &PokemonInfos, nil
}
