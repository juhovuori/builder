package server

import (
	"net/http"

	"github.com/juhovuori/builder/app"
	"github.com/juhovuori/builder/project"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Server is the HTTP server that implements Builder API
type Server interface {
	Run() error
}

type echoServer struct {
	echo  *echo.Echo
	app   app.App
	token *string
}

func (s echoServer) Run() error {
	return s.echo.Start(s.app.Config().ServerAddress)
}

func (s echoServer) setupRouteHandlers() error {
	s.echo.GET("/", s.hRoot)
	s.echo.GET("/health", s.hHealth)
	s.echo.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	s.echo.POST("/v1/builds", s.hCreateBuild)
	s.echo.GET("/v1/builds", s.hListBuilds)
	s.echo.GET("/v1/builds/:id", s.hGetBuild)
	s.echo.GET("/v1/builds/:id/stdout", s.hGetStdout)
	s.echo.POST("/v1/builds/:id", s.hAddStage)
	s.echo.GET("/v1/builds/:id/data/:stage", s.hGetStageData)
	s.echo.GET("/v1/projects", s.hListProjects)
	s.echo.GET("/v1/projects/:id", s.hGetProject)
	s.echo.GET("/v1/projects/:id/builds", s.hListProjectBuilds)
	s.echo.POST("/v1/projects/:id/trigger", s.hTriggerBuild)
	s.echo.GET("/v1/version", s.hVersion)
	s.echo.POST("/v1/shutdown", s.hShutdown)
	return nil
}

func (s echoServer) errorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := err.Error()
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	} else if err == project.ErrNotFound {
		code = 404
	}
	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD { // Issue #608
			c.NoContent(code)
		} else {
			se := serverError{code, msg}
			c.JSON(code, se)
		}
	}
	s.echo.Logger.Error(err)
}

// New creates a new echo Server instance
func New(app app.App) (Server, error) {
	return NewWithSystemToken(app, nil)
}

// NewWithSystemToken creates a new echo Server instance
func NewWithSystemToken(app app.App, token *string) (Server, error) {
	server := echoServer{
		echo:  echo.New(),
		app:   app,
		token: token,
	}
	server.echo.HTTPErrorHandler = server.errorHandler
	server.echo.Pre(middleware.RemoveTrailingSlash())
	server.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	server.setupRouteHandlers()
	return server, nil
}

type serverError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
