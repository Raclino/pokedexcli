package main

import (
	"fmt"
	"net/http"
	"time"
)

const limit = 20

var client = &http.Client{Timeout: 3 * time.Second}
var startRange = 1

func commandMap(config *Config) error {
	fmt.Println(*config)
	endRange := startRange + limit - 1
	FetchLocationsNames(client, startRange, endRange)
	startRange += limit
	return nil
}

func commandMapb(config *Config) error {
	fmt.Println(*config)
	startRange -= limit * 2
	endRange := startRange + limit - 1
	FetchLocationsNames(client, startRange, endRange)
	return nil
}
