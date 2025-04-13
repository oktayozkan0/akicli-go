package api

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type BuildAppRequest struct {
	Tag        string `json:"tag"`
	PatchNotes string `json:"patch_notes"`
}

func (a *API) GetApps(params map[string]string) (*PaginatedResponse[Application], error) {
	return FetchResource[PaginatedResponse[Application]](a.client, ApplicationsPath, params)
}

func (a *API) GetApp(appid int) (*Application, error) {
	u := fmt.Sprintf(ApplicationDetailPath, appid)
	return FetchResource[Application](a.client, u, nil)
}

func (a *API) GetAppVersions(appid int, params map[string]string) (*PaginatedResponse[ApplicationVersion], error) {
	u := fmt.Sprintf(ApplicationVersionsPath, appid)
	return FetchResource[PaginatedResponse[ApplicationVersion]](a.client, u, params)
}

func (a *API) GetAppVersionDetails(appid, versionid int) (*ApplicationVersion, error) {
	u := fmt.Sprintf(ApplicationVersionDetailPath, appid, versionid)
	return FetchResource[ApplicationVersion](a.client, u, nil)
}

func (a *API) BuildApp(appid int, tag, notes string) error {
	u := fmt.Sprintf(ApplicationBuildPath, appid)
	buildPayload := BuildAppRequest{
		Tag:        tag,
		PatchNotes: notes,
	}
	requestData, err := json.Marshal(buildPayload)
	if err != nil {
		return err
	}
	_, err = a.client.Post(u, bytes.NewBuffer(requestData))
	if err != nil {
		return err
	}
	return nil
}
