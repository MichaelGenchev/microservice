package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MichaelGenchev/microservice/internal/businesslogic"
)
type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{endpoint: endpoint}
}

func (c *Client) GetWorkout(id string) (*businesslogic.Workout, error) {
	endpoint := c.endpoint + "/" + id
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		errResp := map[string]any{}
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("service responded with a non 200 status code: %s", errResp["error"])
	}
	workoutResp := new(businesslogic.Workout)
	if err := json.NewDecoder(resp.Body).Decode(workoutResp); err != nil {
		return nil, err
	}
	resp.Body.Close()

	return workoutResp, nil
}