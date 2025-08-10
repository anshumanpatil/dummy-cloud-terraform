package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL is the default URL for the bakwas-cloud API.
const HostURL string = "http://localhost:8080"

// Client is the main struct for interacting with the bakwas-cloud API.
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

type AuthStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func NewClient(host, username, password *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// If username or password not provided, return empty client
	if username == nil || password == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Email:    *username,
		Password: *password,
	}

	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("Authorization", token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
