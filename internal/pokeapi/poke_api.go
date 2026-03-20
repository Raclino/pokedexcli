package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func GetLocationAreas(client *http.Client, url string) (*LocationAreasResponse, error) {
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

	locationAreasResponse := LocationAreasResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&locationAreasResponse); err != nil {
		return nil, fmt.Errorf("couldn't decode response body: %w", err)
	}

	return &locationAreasResponse, nil
}
