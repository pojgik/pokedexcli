package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pojgik/pokedexcli/internal/pokecache"
)

func ListLocations(pageURL *string, locationsCache *pokecache.Cache) (locationList, error) {
	var url string
	if pageURL != nil {
		url = *pageURL
	} else {
		url = "https://pokeapi.co/api/v2/location-area"
	} // if

	data, ok := locationsCache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return locationList{}, err
		} // if
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return locationList{}, fmt.Errorf("response failed with status code: %d and \nbody: %s", res.StatusCode, body)
		} // if
		if err != nil {
			return locationList{}, err
		} // if
		locationsCache.Add(url, body)
		data = body
	} else {
		fmt.Println("Accessing data from cache")
	} // if

	locationsList := locationList{}
	err := json.Unmarshal(data, &locationsList)
	if err != nil {
		return locationList{}, err
	} // if

	return locationsList, nil

}
