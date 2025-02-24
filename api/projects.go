package api

type Project struct {
	Pk                int    `json:"pk"`
	Name              string `json:"name"`
	Slug              string `json:"slug"`
	Account           int    `json:"account"`
	TotalAppCount     int    `json:"total_app_count"`
	TotalServiceCount int    `json:"total_service_count"`
}

type ProjectPaginated struct {
	PaginatedResponse
	Results []Project `json:"results"`
}

func (a *API) GetProjects(params map[string]string) (*ProjectPaginated, error) {
	u := projectsPath
	return FetchResource[ProjectPaginated](a.client, u, params)
}
