package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pojgik/pokedexcli/internal/pokecache"
)

func Explore(url string, encountersCache *pokecache.Cache) (LocationDetails, error) {

	data, ok := encountersCache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return LocationDetails{}, err
		} // if
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return LocationDetails{}, fmt.Errorf("response failed with status code: %d and \nbody: %s", res.StatusCode, body)
		} // if
		if err != nil {
			return LocationDetails{}, err
		} // if
		encountersCache.Add(url, body)
		data = body
	} else {
		fmt.Println("Accessing data from cache")
	} // if

	locationDetails := LocationDetails{}

	err := json.Unmarshal(data, &locationDetails)
	if err != nil {
		return LocationDetails{}, err
	} // if

	return locationDetails, nil

}
