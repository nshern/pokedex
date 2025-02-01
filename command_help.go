package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	cfg := &Config{
		Next:     "",
		Previous: "",
	}
	for _, command := range getCommands(cfg) {
		fmt.Println()
		fmt.Printf("%s: %s", command.name, command.description)
	}
	fmt.Println()
	return nil

}
