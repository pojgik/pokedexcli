package main

import (
	"fmt"

	"github.com/pojgik/pokedexcli/internal/pokeapi"
)

func commandExplore(config *config, param string, caught map[string]pokeapi.Pokemon) error {
	url := "https://pokeapi.co/api/v2/location-area/" + param
	locationDetails, err := pokeapi.Explore(url, &config.cache)
	if err != nil {
		return err
	} // if
	fmt.Printf("Exploring %s...\n", param)
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationDetails.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	} // for
	return nil
} // explore
