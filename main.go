package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		newInputText := scanner.Text()
		cleanedInput := cleanInput(newInputText)

		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]

		c, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := c.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}
