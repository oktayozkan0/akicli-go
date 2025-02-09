package api

import (
	"encoding/json"

	"github.com/oktayozkan0/akicli-go/utils"
)

type Application struct {
	Pk   int    `json:"pk"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Applications struct {
	PaginatedResponse
	Results []Application
}

func (a *API) GetApps(params map[string]string) (*Applications, error) {
	apps, err := a.client.Get(applicationsPath, params)
	if err != nil {
		return nil, err
	}
	body, err := utils.ResponseAsBytes(apps)
	if err != nil {
		return nil, err
	}
	var response Applications
	err = json.Unmarshal(body, &response)
	return &response, err
}
