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
