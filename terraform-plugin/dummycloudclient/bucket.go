package dummycloudclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetBuckets - Returns list of Buckets.
func (c *Client) GetBuckets(id string) ([]Bucket, error) {
	i := BucketRead{
		ID: id,
	}
	readBody, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/bucket/read", c.HostURL), strings.NewReader(string(readBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	coffees := []Bucket{}
	err = json.Unmarshal(body, &coffees)
	if err != nil {
		return nil, err
	}

	return coffees, nil
}

// CreateBucket - Returns list of Buckets.
func (c *Client) CreateBucket(i Bucket) (*Bucket, error) {
	createBody, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/bucket", c.HostURL), strings.NewReader(string(createBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	coffees := Bucket{}
	err = json.Unmarshal(body, &coffees)
	if err != nil {
		return nil, err
	}

	return &coffees, nil
}

// UpdateBucket - Returns list of Buckets.
func (c *Client) UpdateBucket(i Bucket) (*Bucket, error) {
	updateBody, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/bucket", c.HostURL), strings.NewReader(string(updateBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	coffees := Bucket{}
	err = json.Unmarshal(body, &coffees)
	if err != nil {
		return nil, err
	}

	return &coffees, nil
}

// DeleteBucket - Returns list of Buckets.
func (c *Client) DeleteBucket(id string) (*DeleteBucket, error) {
	i := BucketDelete{
		ID: id,
	}
	readBody, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/bucket", c.HostURL), strings.NewReader(string(readBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	coffees := DeleteBucket{}
	err = json.Unmarshal(body, &coffees)
	if err != nil {
		return nil, err
	}

	return &coffees, nil
}
