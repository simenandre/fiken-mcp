package fiken

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const baseURL = "https://api.fiken.no/api/v2"

// Client is an HTTP client for the Fiken API.
type Client struct {
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new Fiken API client.
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

// Do executes an HTTP request against the Fiken API.
// Returns (body, statusCode, error).
func (c *Client) Do(method, path string, body []byte, queryParams map[string]string) ([]byte, int, error) {
	u, err := url.Parse(baseURL + path)
	if err != nil {
		return nil, 0, fmt.Errorf("invalid URL: %w", err)
	}

	if len(queryParams) > 0 {
		q := u.Query()
		for k, v := range queryParams {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return nil, 0, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("reading response body: %w", err)
	}

	return respBody, resp.StatusCode, nil
}

// Get performs a GET request.
func (c *Client) Get(path string, queryParams map[string]string) ([]byte, int, error) {
	return c.Do(http.MethodGet, path, nil, queryParams)
}

// Post performs a POST request.
func (c *Client) Post(path string, body []byte) ([]byte, int, error) {
	return c.Do(http.MethodPost, path, body, nil)
}

// Put performs a PUT request.
func (c *Client) Put(path string, body []byte) ([]byte, int, error) {
	return c.Do(http.MethodPut, path, body, nil)
}

// Delete performs a DELETE request.
func (c *Client) Delete(path string) ([]byte, int, error) {
	return c.Do(http.MethodDelete, path, nil, nil)
}

// BuildQueryParams builds query params from key-value pairs.
// Nil values and empty strings are omitted.
func BuildQueryParams(pairs ...interface{}) map[string]string {
	params := make(map[string]string)
	for i := 0; i+1 < len(pairs); i += 2 {
		key, ok := pairs[i].(string)
		if !ok {
			continue
		}
		val := pairs[i+1]
		if val == nil {
			continue
		}
		switch v := val.(type) {
		case string:
			if v != "" {
				params[key] = v
			}
		case int:
			params[key] = strconv.Itoa(v)
		case float64:
			params[key] = strconv.FormatFloat(v, 'f', -1, 64)
		}
	}
	return params
}
