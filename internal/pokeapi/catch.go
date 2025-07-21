package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pojgik/pokedexcli/internal/pokecache"
)

func Catch(url string, pokemonCache *pokecache.Cache) (Pokemon, error) {
	data, ok := pokemonCache.Get(url)
	if ok {
		fmt.Println("Accessing data from cache")
	} else {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		} // if
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return Pokemon{}, fmt.Errorf("response failed with status code: %d and \nbody: %s", res.StatusCode, body)
		} // if
		if err != nil {
			return Pokemon{}, err
		} // if
		pokemonCache.Add(url, body)
		data = body
	} // if

	pokemon := Pokemon{}

	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	} // if

	return pokemon, nil
} // Catch
