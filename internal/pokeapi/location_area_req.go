package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/harsh-m-patil/pokedex/internal/pokeapi/types"
)

func (c *Client) ListLocationAreas() (types.LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return types.LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return types.LocationAreaResp{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return types.LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.LocationAreaResp{}, err
	}

	locationAreasResp := types.LocationAreaResp{}

	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return types.LocationAreaResp{}, err
	}

	return locationAreasResp, nil
}
