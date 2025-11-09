package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) GetPokemonStats(pokemonURL string) (RespPokemon, error) {

	if val, ok := c.cache.Get(pokemonURL); ok {
		pokemonStats := RespPokemon{}
		if err := json.Unmarshal(val, &pokemonStats); err != nil {
			return RespPokemon{}, err
		}

		return pokemonStats, nil
	}

	res, err := c.httpClient.Get(pokemonURL)
	if err != nil {
		return RespPokemon{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespPokemon{}, err
	}
	defer res.Body.Close()

	pokemonStats := RespPokemon{}
	if err := json.Unmarshal(body, &pokemonStats); err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(pokemonURL, body)
	return pokemonStats, nil
}
