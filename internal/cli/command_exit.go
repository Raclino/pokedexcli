package cli

import (
	"fmt"
	"os"
)

func commandExit(appConfig *AppConfig, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
