package client

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	baseURL    string
	Token      string
	httpClient *http.Client
}

type ClientOption func(*Client)

func WithBaseURL(baseUrl string) ClientOption {
	return func(c *Client) {
		u, err := url.Parse(baseUrl)
		if err != nil {
			panic(err)
		}
		hasSlashSuffix := strings.HasSuffix(u.String(), "/")
		if hasSlashSuffix {
			u.Path = u.Path[:len(u.Path)-1]
		}
		pathHasApiSuffix := strings.HasSuffix(u.Path, "api/v1")
		if !pathHasApiSuffix {
			u = u.JoinPath("api", "v1")
		}
		c.baseURL = u.String()
	}
}

func WithToken(token string) ClientOption {
	return func(c *Client) {
		c.Token = token
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
	if c.Token != "" {
		req.Header.Set("Authorization", c.Token)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "akinoncli go")
	return c.httpClient.Do(req)
}

func (c *Client) Get(path string, params map[string]string) (*http.Response, error) {
	u := c.baseURL + path
	req, err := http.NewRequest("GET", u, nil)
	queries := req.URL.Query()
	for k, v := range params {
		queries.Add(k, v)
	}
	req.URL.RawQuery = queries.Encode()
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) Post(path string, body *bytes.Buffer) (*http.Response, error) {
	u := c.baseURL + path
	req, err := http.NewRequest("POST", u, body)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}
