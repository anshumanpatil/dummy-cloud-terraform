package dummycloudclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetNetworks - Returns list of Networks.
func (c *Client) GetNetworks(id string) ([]Network, error) {
	i := NetworkRead{
		ID: id,
	}
	readBody, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/network/read", c.HostURL), strings.NewReader(string(readBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	coffees := []Network{}
	err = json.Unmarshal(body, &coffees)
	if err != nil {
		return nil, err
	}

	return coffees, nil
}

// CreateNetwork - Returns list of Networks.
func (c *Client) CreateNetwork(i Network) (*Network, error) {
	createBody, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/network", c.HostURL), strings.NewReader(string(createBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	coffees := Network{}
	err = json.Unmarshal(body, &coffees)
	if err != nil {
		return nil, err
	}

	return &coffees, nil
}

// UpdateNetwork - Returns list of Networks.
func (c *Client) UpdateNetwork(i Network) (*Network, error) {
	updateBody, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/network", c.HostURL), strings.NewReader(string(updateBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	coffees := Network{}
	err = json.Unmarshal(body, &coffees)
	if err != nil {
		return nil, err
	}

	return &coffees, nil
}

// DeleteNetwork - Returns list of Networks.
func (c *Client) DeleteNetwork(id string) (*DeleteNetwork, error) {
	i := NetworkDelete{
		ID: id,
	}
	readBody, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/network", c.HostURL), strings.NewReader(string(readBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	coffees := DeleteNetwork{}
	err = json.Unmarshal(body, &coffees)
	if err != nil {
		return nil, err
	}

	return &coffees, nil
}
