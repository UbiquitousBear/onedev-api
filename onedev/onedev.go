package onedev

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	BaseURL   *url.URL
	client    http.Client
	UserAgent string

	Projects *ProjectService
}

type service struct {
	client *Client
}

type Response struct {
	StatusCode int             `json:"code"`
	Data       json.RawMessage `json:"data"`
}

func NewClient(baseUrl string) (*Client, error) {
	baseEndpoint, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(baseEndpoint.Path, "/") {
		baseEndpoint.Path += "/"
	}

	return &Client{client: http.Client{}, BaseURL: baseEndpoint}, nil
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	response, err := parseResponse(resp)
	if err != nil {
		return nil, err
	}

	switch v := v.(type) {
	case nil:
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}

	return response, nil
}

// parseResponse parses a successful response
func parseResponse(res *http.Response) (*Response, error) {
	defer res.Body.Close()
	response := &Response{
		StatusCode: res.StatusCode,
	}

	if err := json.NewDecoder(res.Body).Decode(response); err != nil && err != io.EOF { // io.EOF means an empty body
		return nil, err
	}
	log.Printf("[DEBUG] received response with body %s", res.Body)
	return response, nil
}
