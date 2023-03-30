package api

import (
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) initRoutes() {
	s.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := s.Group("/api/v1")
	{
		// Task related routes
		v1.POST("/task", createTask(s.ss.taskSvc))
		v1.GET("/task/:id", GetTask(s.ss.taskSvc))
	}
}
