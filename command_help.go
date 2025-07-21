package main

import (
	"fmt"

	"github.com/pojgik/pokedexcli/internal/pokeapi"
)

func commandHelp(config *config, param string, caught map[string]pokeapi.Pokemon) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	} // for
	fmt.Println()
	return nil
} // commandHelp
