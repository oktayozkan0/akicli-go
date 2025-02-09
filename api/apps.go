package api

import (
	"encoding/json"
	"strconv"

	"github.com/oktayozkan0/akicli-go/utils"
)

type ApplicationType struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Application struct {
	Pk              int             `json:"pk"`
	Name            string          `json:"name"`
	Slug            string          `json:"slug"`
	AppType         string          `json:"app_type"`
	ApplicationType ApplicationType `json:"application_type"`
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

func (a *API) GetApp(id int) (*Application, error) {
	u := "/applications/" + strconv.Itoa(id) + "/"
	app, err := a.client.Get(u, nil)
	if err != nil {
		return nil, err
	}
	body, err := utils.ResponseAsBytes(app)
	if err != nil {
		return nil, err
	}
	var response Application
	err = json.Unmarshal(body, &response)
	return &response, err
}
