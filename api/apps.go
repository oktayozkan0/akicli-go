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

type ApplicationBuildStatus struct {
}

type ApplicationVersion struct {
	Pk         int    `json:"pk"`
	App        string `json:"app"`
	Tag        string `json:"tag"`
	Status     string `json:"status"`
	PatchNotes string `json:"patch_notes"`
}

func (a *API) GetApps(params map[string]string) (*PaginatedResponse[Application], error) {
	return FetchResource[PaginatedResponse[Application]](a.client, applicationsPath, params)
}

func (a *API) GetApp(id int) (*Application, error) {
	u := applicationsPath + "/" + strconv.Itoa(id) + "/"
	return FetchResource[Application](a.client, u, nil)
}

func (a *API) GetAppVersions(id int, params map[string]string) (*PaginatedResponse[ApplicationVersion], error) {
	u := applicationsPath + "/" + strconv.Itoa(id) + "/versions/"
	return FetchResource[PaginatedResponse[ApplicationVersion]](a.client, u, params)
}

// func (a *API) BuildApp(id int) (*ApplicationBuildStatus, error) {

// }
