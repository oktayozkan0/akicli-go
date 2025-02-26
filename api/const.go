package api

const (
	ApplicationsPath             = "/applications"                      // implemented
	ApplicationDetailPath        = "/applications/%d/"                  // implemented
	ApplicationBuildPath         = "/applications/%d/build/"            // implemented
	ApplicationVersionsPath      = "/applications/%d/versions/"         // implemented
	ApplicationVersionDetailPath = "/applications/%d/versions/%d/"      // implemented
	ApplicationVersionLogsPath   = "/applications/%d/versions/%d/logs/" //notfound on server
	ProjectsPath                 = "/projects/"                         // implemented
	ProjectDetailPath            = "/projects/%d/"
	ProjectAppsPath              = "/projects/%d/project_apps/"
	ProjectAppDetailPath         = "/projects/%d/project_apps/%d/"
	ProjectAppCustomEnvPath      = "/projects/%d/project_apps/%d/custom_env/"
	ProjectAppDeployPath         = "/projects/%d/project_apps/%d/deploy/"
	ProjectAppDeployMultiplePath = "/projects/%d/project_apps/%d/deploy_multiple_project_apps/"
	ProjectAppDeploymentsPath    = "/projects/%d/project_apps/%d/deployments/"
	ProjectAppDeploymentLogPath  = "/projects/%d/project_apps/%d/deployments/%d/logs/"
	LoginPath                    = "/users/login/" // implemented
)
