package main

import "strings"

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	loweredText := strings.ToLower(trimmedText)

	inputCleaned := []string{}

	for w := range strings.FieldsSeq(loweredText) {
		inputCleaned = append(inputCleaned, w)
	}

	return inputCleaned
}
