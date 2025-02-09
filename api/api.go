package api

import (
	"github.com/oktayozkan0/akicli-go/client"
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
