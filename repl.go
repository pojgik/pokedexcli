package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pojgik/pokedexcli/internal/pokeapi"
	"github.com/pojgik/pokedexcli/internal/pokecache"
)

type config struct {
	Next     *string
	Previous *string
	cache    pokecache.Cache
	caught   map[string]pokeapi.Pokemon
}

func startRepl() {
	config := &config{}
	config.cache = *pokecache.NewCache(5 * time.Second)
	config.caught = make(map[string]pokeapi.Pokemon)
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
		var param string
		if len(userInput) > 1 {
			param = userInput[1]
		} // if

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, param)
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
	callback    func(*config, string) error
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
		"explore": {
			name:        "explore",
			description: "Explore an area, listing all of the potential encounters there",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch the specified Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Check information about a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List the names of all pokemon the user has caught",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	} // cliCommand
} // getCommands
