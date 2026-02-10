package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	body, exists := c.cache.Get(url)

	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		} 
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		} 
		if resp.StatusCode > 299 {
			return Pokemon{}, fmt.Errorf("unable to complete request. Status code: %d", resp.StatusCode)
		}
		body, err = io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return Pokemon{}, err
		} 
		c.cache.Add(url, body)
	} 

	var data Pokemon
	if err := json.Unmarshal(body, &data); err != nil {
		return Pokemon{}, err
	} 
	return data, nil
} 
