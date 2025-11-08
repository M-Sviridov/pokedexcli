package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) LocationAreaList(pageURL *string) (RespLocationArea, error) {
	url := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locations := RespLocationArea{}
		if err := json.Unmarshal(val, &locations); err != nil {
			return RespLocationArea{}, err
		}

		return locations, nil
	}

	res, err := c.httpClient.Get(url)
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

	c.cache.Add(url, body)
	return locations, nil
}
