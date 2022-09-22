package item

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GET items
func (c *Client) GetItems() ([]Item, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/items", c.HostURL), nil)

	if err != nil {
		return nil, err
	}

	rb, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var items []Item

	err = json.Unmarshal(rb, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// GET item by title
func (c *Client) GetItem(title string) (*Item, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/item?title=%s", c.HostURL, title), nil)
	if err != nil {
		return nil, err
	}

	rb, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var item Item

	err = json.Unmarshal(rb, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// POST item
func (c *Client) CreateItem(item Item) (*Item, error) {
	rb, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/item", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var finalItem Item
	err = json.Unmarshal(body, &finalItem)
	if err != nil {
		return nil, err
	}

	return &finalItem, nil
}

// PUT item
func (c *Client) UpdateItem(item Item) (*Item, error) {
	rb, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/item", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var finalItem Item
	err = json.Unmarshal(body, &finalItem)
	if err != nil {
		return nil, err
	}

	return &finalItem, nil
}

// DELETE item
func (c *Client) DeleteItem(item Item) (*Item, error) {
	rb, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/item", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, nil
	}

	body, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var finalItem Item
	err = json.Unmarshal(body, &finalItem)
	if err != nil {
		return nil, err
	}

	return &finalItem, nil
}
