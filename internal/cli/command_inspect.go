package cli

import (
	"fmt"
	"strings"

	"github.com/Raclino/pokedexcli/internal/pokeapi"
)

func commandInspect(config *pokeapi.LocationAreaConfig, args ...string) error {
	pokemonName := args[1]

	p, ok := Pokedex[pokemonName]
	if !ok {
		fmt.Printf("%s is not in your Pokedex, catch him first for more informations\n", pokemonName)
		return nil
	}

	fmt.Print(formatPokemonInfos(p))
	return nil
}

func formatPokemonInfos(p pokeapi.PokemonInfos) string {
	var b strings.Builder

	fmt.Fprintf(&b, "Name: %s\n", p.Name)
	fmt.Fprintf(&b, "Height: %d\n", p.Height)
	fmt.Fprintf(&b, "Weight: %d\n", p.Weight)

	b.WriteString("Stats:\n")
	for _, stat := range p.Stats {
		fmt.Fprintf(&b, "  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	b.WriteString("Types:\n")
	for _, t := range p.Types {
		fmt.Fprintf(&b, "  - %s\n", t.Type.Name)
	}

	return b.String()
}
