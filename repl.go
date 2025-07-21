package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pojgik/pokedexcli/internal/pokecache"
)

type config struct {
	Next           *string
	Previous       *string
	locationsCache pokecache.Cache
}

func startRepl() {
	config := &config{}
	config.locationsCache = *pokecache.NewCache(5 * time.Second)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		success := scanner.Scan()
		if !success {
			err := scanner.Err()
			if err != nil {
				fmt.Printf("Error scanning input: %v\n", err)
			} // if
		} // if
		userInput := cleanInput(scanner.Text())
		commandName := userInput[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			} // if
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		} // else
	} // for
} // startRepl

func cleanInput(text string) []string {
	output := strings.Fields(strings.ToLower(text))
	return output
} // cleanInput

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
} // cliCommand

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	} // cliCommand
} // getCommands
