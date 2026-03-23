package cli

import (
	"fmt"

	"github.com/Raclino/pokedexcli/internal/pokeapi"
)

func updateConfigFromResponse(appConfig *AppConfig, resp *pokeapi.LocationAreasResponse) {
	if resp.Next != nil {
		appConfig.locationAreaUrls.Next = *resp.Next
	} else {
		appConfig.locationAreaUrls.Next = ""
	}

	if resp.Previous != nil {
		appConfig.locationAreaUrls.Previous = *resp.Previous
	} else {
		appConfig.locationAreaUrls.Previous = ""
	}
}

func commandMap(appConfig *AppConfig, args ...string) error {
	resp, err := pokeapi.GetLocationAreas(appConfig.client, &appConfig.cache, appConfig.locationAreaUrls.Next)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	updateConfigFromResponse(appConfig, resp)
	return nil
}

func commandMapb(appConfig *AppConfig, args ...string) error {
	if appConfig.locationAreaUrls.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	resp, err := pokeapi.GetLocationAreas(appConfig.client, &appConfig.cache, appConfig.locationAreaUrls.Previous)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	updateConfigFromResponse(appConfig, resp)
	return nil
}
