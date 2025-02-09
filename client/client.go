package client

import (
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

func NewClient(options ...ClientOption) *Client {
	client := &Client{
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
	for _, option := range options {
		option(client)
	}
	return client
}
