package api

import (
	"strconv"
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
	return FetchResource[ApplicationPaginated](a.client, applicationsPath, params)
}

func (a *API) GetApp(id int) (*Application, error) {
	u := applicationsPath + "/" + strconv.Itoa(id) + "/"
	return FetchResource[Application](a.client, u, nil)
}

func (a *API) GetAppVersions(id int, params map[string]string) (*ApplicationVersionPaginated, error) {
	u := applicationsPath + "/" + strconv.Itoa(id) + "/versions/"
	return FetchResource[ApplicationVersionPaginated](a.client, u, params)
}

// func (a *API) BuildApp(id int) (*ApplicationBuildStatus, error) {

// }
