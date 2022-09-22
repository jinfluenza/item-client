package item

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const HostUrl string = "http://localhost:4040"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
}

func NewClient() (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostUrl,
	}

	return &c, nil
}

func (c *Client) sendRequest(rq *http.Request) ([]byte, error) {
	res, err := c.HTTPClient.Do(rq)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
