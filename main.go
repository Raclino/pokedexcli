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

		fmt.Println("Your command was: ", cleanedInput[0])
		continue
	}
}
