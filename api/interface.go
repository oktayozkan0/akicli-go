package api

type APIInterface interface {
	Login(email, password string) error
	GetProjects(params map[string]string) (*PaginatedResponse[Project], error)
	GetApps(params map[string]string) (*PaginatedResponse[Application], error)
	GetApp(appid int) (*Application, error)
	GetAppVersions(appid int, params map[string]string) (*PaginatedResponse[ApplicationVersion], error)
	GetAppVersionDetails(appid, versionid int) (*ApplicationVersion, error)
	BuildApp(appid int, tag string) error
}
