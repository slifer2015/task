package api

import (
	"project/internal/config"
	"project/internal/services/generator"
	"project/internal/services/task"
	"project/internal/services/task/requester"
	"project/internal/services/task/storage/memomry"
	"project/pkg/logger"
)

type ServiceStorage struct {
	logger  *logger.Logger
	taskSvc taskService
}

const (
	jobBufferSize = 1024
)

func newServiceStorage(cfg *config.Config, logger *logger.Logger) *ServiceStorage {
	out := &ServiceStorage{}
	out.logger = logger
	storage := memomry.NewMemory()
	requestHandler := requester.NewRequester()
	out.taskSvc = task.NewService(
		logger,
		storage,
		requestHandler,
		generator.NewGenerator(),
		cfg.WorkersCount,
		jobBufferSize,
	)
	return out
}
