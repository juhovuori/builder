package server

import "github.com/labstack/echo"

// Server is the HTTP server that implements Builder API
type Server interface {
	Run() error
}

type echoServer struct {
	echo *echo.Echo
	cfg  Config
}

func (s *echoServer) Run() error {
	return s.echo.Start(s.cfg.ServerAddress)
}

func (s *echoServer) setupRouteHandlers() error {
	s.echo.GET("/", s.hRoot)
	s.echo.GET("/health", s.hHealth)
	s.echo.POST("/v1/builds", s.hCreateBuild)
	s.echo.GET("/v1/builds", s.hListBuilds)
	s.echo.GET("/v1/builds/:id", s.hGetBuild)
	s.echo.POST("/v1/builds/:id", s.hRouteStage)
	s.echo.GET("/v1/projects", s.hListProjects)
	s.echo.GET("/v1/projects/:id", s.hGetProject)
	return nil
}

// New creates a new echo Server instance
func New(cfg Config) (Server, error) {
	echo := echo.New()
	server := echoServer{
		echo,
		cfg,
	}
	server.setupRouteHandlers()
	return &server, nil
}
