package cli

import (
	"net/http"
	"time"

	"github.com/Raclino/pokedexcli/internal/pokeapi"
	"github.com/Raclino/pokedexcli/internal/pokecache"
)

type AppConfig struct {
	cache            pokecache.Cache
	client           *http.Client
	locationAreaUrls *pokeapi.LocationAreaUrls
	pokedex          map[string]pokeapi.PokemonInfos
}

func NewAppConfig(timeout, cacheInterval time.Duration) *AppConfig {
	return &AppConfig{
		cache:  *pokecache.NewCache(cacheInterval),
		client: &http.Client{Timeout: timeout},
		locationAreaUrls: &pokeapi.LocationAreaUrls{
			Next:     pokeapi.LocationsAreas,
			Previous: "",
		},
		pokedex: make(map[string]pokeapi.PokemonInfos),
	}
}
