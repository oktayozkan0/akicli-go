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

type ApplicationPaginated struct {
	PaginatedResponse
	Results []Application `json:"results"`
}

type ApplicationBuildStatus struct {
}

type ApplicationVersion struct {
	Pk         int    `json:"pk"`
	App        string `json:"app"`
	Tag        string `json:"tag"`
	Status     string `json:"status"`
	PatchNotes string `json:"patch_notes"`
}

type ApplicationVersionPaginated struct {
	PaginatedResponse
	Results []ApplicationVersion `json:"results"`
}

func (a *API) GetApps(params map[string]string) (*ApplicationPaginated, error) {
	apps, err := a.client.Get(applicationsPath, params)
	if err != nil {
		return nil, err
	}
	body, err := utils.ResponseAsBytes(apps)
	if err != nil {
		return nil, err
	}
	var response ApplicationPaginated
	err = json.Unmarshal(body, &response)
	return &response, err
}

func (a *API) GetApp(id int) (*Application, error) {
	u := applicationsPath + "/" + strconv.Itoa(id) + "/"
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

func (a *API) GetAppVersions(id int, params map[string]string) (*ApplicationVersionPaginated, error) {
	u := applicationsPath + "/" + strconv.Itoa(id) + "/versions/"
	app, err := a.client.Get(u, params)
	if err != nil {
		return nil, err
	}
	body, err := utils.ResponseAsBytes(app)
	if err != nil {
		return nil, err
	}
	var response ApplicationVersionPaginated
	err = json.Unmarshal(body, &response)
	return &response, err
}

// func (a *API) BuildApp(id int) (*ApplicationBuildStatus, error) {

// }
