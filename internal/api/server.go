package api

import (
	"project/internal/config"
	"project/pkg/logger"

	"github.com/labstack/echo/v4"
)

type Server struct {
	*echo.Echo
	cfg *config.Config
	ss  *ServiceStorage
}

func NewServer(
	cfg *config.Config,
	log *logger.Logger) (*Server, error) {
	s := &Server{
		cfg:  cfg,
		Echo: echo.New(),
	}

	s.ss = newServiceStorage(cfg, log)

	// run task workers
	s.ss.taskSvc.Run()

	s.initRoutes()

	return s, nil
}

func (s *Server) Close() error {
	_ = s.Echo.Close()
	return nil
}
