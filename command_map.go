package main

import (
	"errors"
	"fmt"

	"github.com/pojgik/pokedexcli/internal/pokeapi"
)

func commandMap(config *config, param string) error {
	locationsList, err := pokeapi.ListLocations(config.Next, &config.cache)
	if err != nil {
		return err
	} // if

	config.Next = locationsList.Next
	config.Previous = locationsList.Previous

	for _, location := range locationsList.Results {
		fmt.Println(location.Name)
	} // for
	return nil
} // commandMap

func commandMapb(config *config, param string) error {
	if config.Previous == nil {
		return errors.New("you're on the first page")
	} // if
	locationsList, err := pokeapi.ListLocations(config.Previous, &config.cache)
	if err != nil {
		return err
	} // if

	config.Next = locationsList.Next
	config.Previous = locationsList.Previous

	for _, location := range locationsList.Results {
		fmt.Println(location.Name)
	} // for
	return nil
} // commandMap
