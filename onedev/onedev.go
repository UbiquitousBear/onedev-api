package onedev

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const userAgent = "go-onedev-api"

type Client struct {
	BaseURL   *url.URL
	client    http.Client
	UserAgent string
	common service
	Projects *ProjectService
}

type service struct {
	client *Client
}

func NewClient(baseUrl string) (*Client, error) {
	baseEndpoint, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(baseEndpoint.Path, "/") {
		baseEndpoint.Path += "/"
	}

	c := &Client{client: http.Client{}, BaseURL: baseEndpoint, UserAgent: userAgent}
	c.common.client = c
	c.Projects = (*ProjectService)(&c.common)

	return c, nil
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

func (c *Client) Do(ctx context.Context, req *http.Request, decodeEntity interface{}) (*http.Response, error) {
	log.Debugf("received request to url %decodeEntity with body %decodeEntity", req.URL.String(), req.Body)
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
	defer resp.Body.Close()

	switch v := decodeEntity.(type) {
	case nil:
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}

		if err != nil {
			return nil, err
		}
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	log.Debugf("received response with status %d, body: %s", resp.StatusCode, string(bodyBytes))

	return resp, nil
}

func Int(v int) *int { return &v }