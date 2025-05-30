package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon

	if val, ok := c.cache.Get(url); ok {
		locationResp := Pokemon{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Pokemon{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	locationResp := Pokemon{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return Pokemon{}, err
	}

	return locationResp, nil
}
