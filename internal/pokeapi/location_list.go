package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func LocationAreaList(pageURL *string) (RespLocationArea, error) {
	url := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	res, err := http.Get(url)
	if err != nil {
		return RespLocationArea{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer res.Body.Close()

	locations := RespLocationArea{}
	if err := json.Unmarshal(body, &locations); err != nil {
		return RespLocationArea{}, err
	}

	return locations, nil
}
