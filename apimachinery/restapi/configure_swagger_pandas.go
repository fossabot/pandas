// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	apimiddleware "github.com/cloustone/pandas/apimachinery/middlewares"
	"github.com/cloustone/pandas/apimachinery/restapi/operations"
	"github.com/cloustone/pandas/apimachinery/restapi/operations/dashboard"
	"github.com/cloustone/pandas/apimachinery/restapi/operations/device"
	"github.com/cloustone/pandas/apimachinery/restapi/operations/logs"
	"github.com/cloustone/pandas/apimachinery/restapi/operations/model"
	"github.com/cloustone/pandas/apimachinery/restapi/operations/project"
	"github.com/cloustone/pandas/apimachinery/restapi/operations/user"
	"github.com/cloustone/pandas/apimachinery/server"

	models "github.com/cloustone/pandas/models"
)

//go:generate swagger generate server --target .. --name  --spec ../swagger.yaml --principal models.Principal --exclude-main

func configureFlags(api *operations.SwaggerPandasAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SwaggerPandasAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.RoleAuthAuth = func(token string, scopes []string) (*models.Principal, error) {
		return nil, errors.NotImplemented("oauth2 bearer auth (roleAuth) has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.UserLoginUserHandler = user.LoginUserHandlerFunc(func(params user.LoginUserParams) middleware.Responder {
		return server.LoginUser(params)
	})
	api.ProjectGetProjectsProjectIDHandler = project.GetProjectsProjectIDHandlerFunc(func(params project.GetProjectsProjectIDParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation project.AddProjectDevice has not yet been implemented")
	})
	api.ProjectAddProjectDeviceHandler = project.AddProjectDeviceHandlerFunc(func(params project.AddProjectDeviceParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation project.AddProjectDevice has not yet been implemented")
	})
	api.ModelCreateModelHandler = model.CreateModelHandlerFunc(func(params model.CreateModelParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation model.CreateModel has not yet been implemented")
	})
	api.ProjectCreateProjectHandler = project.CreateProjectHandlerFunc(func(params project.CreateProjectParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation project.CreateProject has not yet been implemented")
	})
	api.ModelDeleteModelHandler = model.DeleteModelHandlerFunc(func(params model.DeleteModelParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation model.DeleteModel has not yet been implemented")
	})
	api.ProjectDeleteProjectHandler = project.DeleteProjectHandlerFunc(func(params project.DeleteProjectParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation project.DeleteProject has not yet been implemented")
	})
	api.ProjectDeleteProjectDeviceHandler = project.DeleteProjectDeviceHandlerFunc(func(params project.DeleteProjectDeviceParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation project.DeleteProjectDevice has not yet been implemented")
	})
	api.LogsGetDeviceLogHandler = logs.GetDeviceLogHandlerFunc(func(params logs.GetDeviceLogParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation logs.GetDeviceLog has not yet been implemented")
	})
	api.ModelGetModelHandler = model.GetModelHandlerFunc(func(params model.GetModelParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation model.GetModel has not yet been implemented")
	})
	api.ModelGetModelsHandler = model.GetModelsHandlerFunc(func(params model.GetModelsParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation model.GetModels has not yet been implemented")
	})
	api.ProjectGetProjectDeviceHandler = project.GetProjectDeviceHandlerFunc(func(params project.GetProjectDeviceParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation project.GetProjectDevices has not yet been implemented")
	})
	api.ProjectGetProjectDevicesHandler = project.GetProjectDevicesHandlerFunc(func(params project.GetProjectDevicesParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation project.GetProjectDevices has not yet been implemented")
	})
	api.LogsGetProjectLogHandler = logs.GetProjectLogHandlerFunc(func(params logs.GetProjectLogParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation logs.GetProjectLog has not yet been implemented")
	})
	api.ProjectGetProjectSummaryHandler = project.GetProjectSummaryHandlerFunc(func(params project.GetProjectSummaryParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation project.GetProjectSummary has not yet been implemented")
	})
	api.ProjectGetProjectsHandler = project.GetProjectsHandlerFunc(func(params project.GetProjectsParams, principal *models.Principal) middleware.Responder {
		return server.GetProjects(params, principal)
	})
	api.DeviceSendDataToDeviceHandler = device.SendDataToDeviceHandlerFunc(func(params device.SendDataToDeviceParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation device.SendDataToDevice has not yet been implemented")
	})
	api.ModelUpdateModelHandler = model.UpdateModelHandlerFunc(func(params model.UpdateModelParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation model.UpdateModel has not yet been implemented")
	})
	api.ProjectUpdateProjectHandler = project.UpdateProjectHandlerFunc(func(params project.UpdateProjectParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation project.UpdateProject has not yet been implemented")
	})
	api.ProjectUpdateProjectDeviceHandler = project.UpdateProjectDeviceHandlerFunc(func(params project.UpdateProjectDeviceParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation project.UpdateProjectDevice has not yet been implemented")
	})
	api.ProjectUpdateProjectDeviceStatusHandler = project.UpdateProjectDeviceStatusHandlerFunc(func(params project.UpdateProjectDeviceStatusParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation project.UpdateProjectDeviceStatus has not yet been implemented")
	})

	api.DashboardGetDashboardHandler = dashboard.GetDashboardHandlerFunc(func(params dashboard.GetDashboardParams, principal *models.Principal) middleware.Responder {
		return server.GetDashboard(params, principal)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return apimiddleware.Dashboard(
		apimiddleware.Cross(handler),
	)
}
