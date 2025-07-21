package main

import (
	"fmt"
	"os"

	"github.com/pojgik/pokedexcli/internal/pokeapi"
)

func commandExit(cfg *config, param string, caught map[string]pokeapi.Pokemon) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
