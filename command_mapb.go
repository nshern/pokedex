package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(cfg *Config) error {

	type locationAreas struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}

	url := cfg.Previous
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	resp, _ := http.Get(url)

	body, _ := io.ReadAll(resp.Body)
	var foo locationAreas
	json.Unmarshal(body, &foo)

	for _, v := range foo.Results {
		fmt.Println(v.Name)
	}

	cfg.Next = foo.Next
	cfg.Previous = foo.Previous

	return nil

}
