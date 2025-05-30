package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationData(location string) (RespDeepLocation, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		locationResp := RespDeepLocation{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespDeepLocation{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDeepLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDeepLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDeepLocation{}, err
	}

	locationResp := RespDeepLocation{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespDeepLocation{}, err
	}

	return locationResp, nil
}
