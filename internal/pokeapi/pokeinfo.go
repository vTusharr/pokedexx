package pokeapi


import (
	"net/http"
	"encoding/json"
	"io"
	"errors"
	"strings"
	
)


func (c *Client) Getpoke(name *string) (PokeInfo , error){
	
	n := strings.ToLower(strings.TrimSpace(*name))
	if n == "" {
		return PokeInfo{}, errors.New("name cannot be empty")
	}

	payload := baseURL + "/pokemon/" + n + "/"
	if val, ok := c.cache.Get(payload); ok {
		PokeResp := PokeInfo{}
		err := json.Unmarshal(val, &PokeResp)
		if err != nil {
			return PokeInfo{}, err
		}

		return PokeResp, nil
	}
	
	req, err := http.NewRequest("GET", payload, nil)
	if err != nil {
		return PokeInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeInfo{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeInfo{}, err
	}

	PokeResp := PokeInfo{}
	err = json.Unmarshal(dat, &PokeResp)
	if err != nil {
		return PokeInfo{}, err
	}

	c.cache.Add(payload, dat)
	return PokeResp, nil
	
	
	
}