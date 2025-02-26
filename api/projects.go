package api

type Project struct {
	Pk                int    `json:"pk"`
	Name              string `json:"name"`
	Slug              string `json:"slug"`
	Account           int    `json:"account"`
	TotalAppCount     int    `json:"total_app_count"`
	TotalServiceCount int    `json:"total_service_count"`
}

func (a *API) GetProjects(params map[string]string) (*PaginatedResponse[Project], error) {
	u := ProjectsPath
	return FetchResource[PaginatedResponse[Project]](a.client, u, params)
}
