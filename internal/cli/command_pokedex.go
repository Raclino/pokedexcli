package cli

import (
	"fmt"
)

func commandPokedex(appConfig *AppConfig, args ...string) error {

	fmt.Println("Your Pokedex:")

	for _, p := range appConfig.pokedex {
		fmt.Printf("  - %s\n", p.Name)
	}
	return nil
}
