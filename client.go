package item

import (
	"fmt"
	"io/ioutil"
	"log"
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
		log.Fatalf("Request failed due to: %s", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal("Unable to read the response body due to followin reason: \n", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Status: %d, body: %s", res.StatusCode, body)
		return nil, fmt.Errorf("Status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
