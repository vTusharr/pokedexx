package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
)

func (c *Client) GetLocation(locationName *string)(Location, error) {
	payload:=baseURL+"/location-area/"+*locationName
	
	if val, ok := c.cache.Get(payload); ok {
		locationsResp := Location{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return Location{}, err
		}

		return locationsResp, nil
	}
	
	req,err:=http.NewRequest("GET",payload,nil)
	
	if err != nil {
		return Location{}, err
	}
	
	resp,err:=c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()
	
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationsResp := Location{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(payload, dat)
	return locationsResp, nil

	
	
}
