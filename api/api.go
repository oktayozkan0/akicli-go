package api

import (
	"encoding/json"

	"github.com/oktayozkan0/akicli-go/client"
	"github.com/oktayozkan0/akicli-go/utils"
)

type API struct {
	client *client.Client
}

func NewAPI(opts ...client.ClientOption) (*API, error) {
	c, err := client.NewClient(opts...)
	if err != nil {
		return nil, err
	}
	return &API{client: c}, nil
}

func FetchResource[T any](client *client.Client, path string, params map[string]string) (*T, error) {
	resp, err := client.Get(path, params)
	if err != nil {
		return nil, err
	}
	body, err := utils.ResponseAsBytes(resp)
	if err != nil {
		return nil, err
	}
	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
