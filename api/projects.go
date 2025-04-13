package api

import "fmt"

type Project struct {
	Pk                int    `json:"pk"`
	Name              string `json:"name"`
	Slug              string `json:"slug"`
	Account           int    `json:"account"`
	TotalAppCount     int    `json:"total_app_count"`
	TotalServiceCount int    `json:"total_service_count"`
}

type ProjectApp struct {
	Pk          int    `json:"pk"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	CreatedDate string `json:"created_date"`
	UpdateDate  string `json:"update_date"`
	Url         string `json:"url"`
}

func (a *API) GetProjects(params map[string]string) (*PaginatedResponse[Project], error) {
	u := ProjectsPath
	return FetchResource[PaginatedResponse[Project]](a.client, u, params)
}

func (a *API) GetProject(pid int) (*Project, error) {
	u := fmt.Sprintf(ProjectDetailPath, pid)
	return FetchResource[Project](a.client, u, nil)
}

func (a *API) GetProjectApps(pid int, params map[string]string) (*PaginatedResponse[ProjectApp], error) {
	u := fmt.Sprintf(ProjectAppsPath, pid)
	return FetchResource[PaginatedResponse[ProjectApp]](a.client, u, params)
}

func (a *API) GetProjectApp(pid, appid int) (*ProjectApp, error) {
	u := fmt.Sprintf(ProjectAppDetailPath, pid, appid)
	return FetchResource[ProjectApp](a.client, u, nil)
}
