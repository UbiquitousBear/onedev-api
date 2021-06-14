package onedev_api

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Client holds all of the information required to connect to a server
type Client struct {
	hostname   string
	user       string
	authToken  string
	httpClient *http.Client
}

func NewClient(hostname string, user string, token string) *Client {
	return &Client{
		hostname:   hostname,
		user:       user,
		authToken:  token,
		httpClient: &http.Client{},
	}
}

func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.user, c.authToken)
	switch method {
	case "GET":
	case "DELETE":
	default:
		req.Header.Add("Content-Type", "application/json")
	}



	log.Printf("[DEBUG] sending request to %s with body %s", req.URL, req.Body)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}
	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s/api/%s", c.hostname, path)
}