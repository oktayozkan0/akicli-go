package api

const (
	applicationsPath             string = "/applications"     // implemented
	applicationDetailPath        string = "/applications/%d/" //implemented
	applicationBuildPath         string = "/applications/%d/build/"
	applicationVersionsPath      string = "/applications/%d/versions/"         //implemented
	applicationVersionDetailPath string = "/applications/%d/versions/%d/"      //implemented
	applicationVersionLogsPath   string = "/applications/%d/versions/%d/logs/" //notfound on server
	projectsPath                 string = "/projects/"
	loginPath                    string = "/users/login/"
)
