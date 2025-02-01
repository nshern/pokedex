package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words

}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type Config struct {
	Next     string
	Previous string
}

func startRepl() {
	cfg := &Config{
		Next:     "",
		Previous: "",
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := cleanInput(scanner.Text())
			if len(userInput) == 0 {
				continue
			}
			commandName := userInput[0]

			command, exists := getCommands(cfg)[commandName]

			if exists {
				err := command.callback()
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else {
				fmt.Println("Unknown command")
				continue
			}

		}
	}

}

func getCommands(cfg *Config) map[string]cliCommand {

	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Get help for the Pokedex",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get maps",
			callback: func() error {
				return commandMap(cfg)
			},
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous maps",
			callback: func() error {
				return commandMapb(cfg)
			},
		},
	}

}
