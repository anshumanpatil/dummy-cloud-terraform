package dummycloudclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetInstances - Returns list of Instances.
func (c *Client) GetInstances(id string) ([]Instance, error) {
	i := InstanceRead{
		ID: id,
	}
	readBody, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/instance/read", c.HostURL), strings.NewReader(string(readBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	coffees := []Instance{}
	err = json.Unmarshal(body, &coffees)
	if err != nil {
		return nil, err
	}

	return coffees, nil
}

// CreateInstance - Returns list of Instances.
func (c *Client) CreateInstance(i Instance) (*Instance, error) {
	createBody, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/instance", c.HostURL), strings.NewReader(string(createBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	coffees := Instance{}
	err = json.Unmarshal(body, &coffees)
	if err != nil {
		return nil, err
	}

	return &coffees, nil
}

// UpdateInstance - Returns list of Instances.
func (c *Client) UpdateInstance(i Instance) (*Instance, error) {
	updateBody, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/instance", c.HostURL), strings.NewReader(string(updateBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	coffees := Instance{}
	err = json.Unmarshal(body, &coffees)
	if err != nil {
		return nil, err
	}

	return &coffees, nil
}

// DeleteInstance - Returns list of Instances.
func (c *Client) DeleteInstance(id string) (*DeleteInstance, error) {
	i := InstanceDelete{
		ID: id,
	}
	readBody, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/instance", c.HostURL), strings.NewReader(string(readBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	coffees := DeleteInstance{}
	err = json.Unmarshal(body, &coffees)
	if err != nil {
		return nil, err
	}

	return &coffees, nil
}
