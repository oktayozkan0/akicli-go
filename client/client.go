package client

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

type ClientOption func(*Client)

func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.baseURL = url
	}
}

func WithToken(token string) ClientOption {
	return func(c *Client) {
		c.token = token
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

func NewClient(options ...ClientOption) (*Client, error) {
	client := &Client{
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
	for _, option := range options {
		option(client)
	}
	if client.baseURL == "" {
		return nil, fmt.Errorf("WithBaseURL required")
	}
	return client, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	if c.token != "" {
		req.Header.Set("Authorization", c.token)
	} else {
		return nil, fmt.Errorf("credentials required, you must login or provide a token")
	}
	return c.httpClient.Do(req)
}

func (c *Client) Get(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}
