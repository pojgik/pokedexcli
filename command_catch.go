package main

import (
	"fmt"
	"math/rand"

	"github.com/pojgik/pokedexcli/internal/pokeapi"
)

func commandCatch(config *config, param string, caught map[string]pokeapi.Pokemon) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + param
	pokemon, err := pokeapi.Catch(url, &config.cache)
	if err != nil {
		return err
	} // if
	fmt.Printf("Throwing a Pokeball at %s...\n", param)
	catch_rate := 0.5 / (float64(pokemon.BaseExperience) / 100.0)
	catch_attempt := rand.Float64()
	if catch_attempt <= catch_rate {
		fmt.Printf("%s was caught!\n", param)
		caught[param] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", param)
	}

	return nil
}
