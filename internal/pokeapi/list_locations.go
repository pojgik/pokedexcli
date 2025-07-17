package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ListLocations(pageURL *string) (locationList, error) {
	var url string
	if pageURL != nil {
		url = *pageURL
	} else {
		url = "https://pokeapi.co/api/v2/location-area"
	} // if

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

	locationsList := locationList{}
	err = json.Unmarshal(body, &locationsList)
	if err != nil {
		return locationList{}, err
	} // if

	return locationsList, nil

}
