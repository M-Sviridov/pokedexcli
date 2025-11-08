package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) LocationPokemonList(areaURL string) (RespLocationPokemon, error) {

	if val, ok := c.cache.Get(areaURL); ok {
		pokemons := RespLocationPokemon{}
		if err := json.Unmarshal(val, &pokemons); err != nil {
			return RespLocationPokemon{}, err
		}

		return pokemons, nil
	}

	res, err := c.httpClient.Get(areaURL)
	if err != nil {
		return RespLocationPokemon{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationPokemon{}, err
	}
	defer res.Body.Close()

	pokemons := RespLocationPokemon{}
	if err := json.Unmarshal(body, &pokemons); err != nil {
		return RespLocationPokemon{}, err
	}

	c.cache.Add(areaURL, body)
	return pokemons, nil
}
