package server

import (
	"net/http"

	"github.com/labstack/echo"
)

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

func (s *echoServer) setupRouting() error {
	s.echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	return nil
}

// New creates a new echo Server instance
func New(cfg Config) (Server, error) {
	echo := echo.New()
	server := echoServer{
		echo,
		cfg,
	}
	server.setupRouting()
	return &server, nil
}
