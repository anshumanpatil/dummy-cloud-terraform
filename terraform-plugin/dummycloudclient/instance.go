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

// CreateOrder - Returns list of Instances.
func (c *Client) CreateOrder(i Instance) (*Instance, error) {
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

// CreateOrder - Returns list of Instances.
func (c *Client) UpdateOrder(i Instance) (*Instance, error) {
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

// CreateOrder - Returns list of Instances.
func (c *Client) DeleteOrder(id string) (*DeleteInstance, error) {
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

// GetCoffee - Returns specific coffee (no auth required)
// func (c *Client) GetCoffee(coffeeID string) ([]Coffee, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/coffees/%s", c.HostURL, coffeeID), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	coffees := []Coffee{}
// 	err = json.Unmarshal(body, &coffees)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return coffees, nil
// }

// GetCoffeeIngredients - Returns list of coffee ingredients (no auth required)
// func (c *Client) GetCoffeeIngredients(coffeeID string) ([]Ingredient, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/coffees/%s/ingredients", c.HostURL, coffeeID), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ingredients := []Ingredient{}
// 	err = json.Unmarshal(body, &ingredients)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ingredients, nil
// }

// CreateCoffee - Create new coffee
// func (c *Client) CreateCoffee(coffee Coffee, authToken *string) (*Coffee, error) {
// 	rb, err := json.Marshal(coffee)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("POST", fmt.Sprintf("%s/coffees", c.HostURL), strings.NewReader(string(rb)))
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, authToken)
// 	if err != nil {
// 		return nil, err
// 	}

// 	newCoffee := Coffee{}
// 	err = json.Unmarshal(body, &newCoffee)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &newCoffee, nil
// }

// // CreateCoffeeIngredient - Create new coffee ingredient
// func (c *Client) CreateCoffeeIngredient(coffee Coffee, ingredient Ingredient, authToken *string) (*Ingredient, error) {
// 	reqBody := struct {
// 		CoffeeID     int    `json:"coffee_id"`
// 		IngredientID int    `json:"ingredient_id"`
// 		Quantity     int    `json:"quantity"`
// 		Unit         string `json:"unit"`
// 	}{
// 		CoffeeID:     coffee.ID,
// 		IngredientID: ingredient.ID,
// 		Quantity:     ingredient.Quantity,
// 		Unit:         ingredient.Unit,
// 	}
// 	rb, err := json.Marshal(reqBody)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("POST", fmt.Sprintf("%s/coffees/%d/ingredients", c.HostURL, coffee.ID), strings.NewReader(string(rb)))
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, authToken)
// 	if err != nil {
// 		return nil, err
// 	}

// 	newIngredient := Ingredient{}
// 	err = json.Unmarshal(body, &newIngredient)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &newIngredient, nil
// }
