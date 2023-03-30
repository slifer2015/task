package task

import (
	"project/pkg/logger"

	"project/internal/common/errorx"
	"project/internal/services/task/requester"
	"project/internal/services/task/storage"
)

type generator interface {
	UniqueID() string
}

type Service struct {
	logger    *logger.Logger
	requester requester.Requester
	jobBuffer chan CreateTaskInput
	storage   storage.Storage
	workers   uint
	generator generator
}

func NewService(
	logger *logger.Logger,
	storage storage.Storage,
	requester requester.Requester,
	generator generator,
	workers uint,
	jobBufferSize int) *Service {
	return &Service{
		logger:    logger,
		requester: requester,
		storage:   storage,
		jobBuffer: make(chan CreateTaskInput, jobBufferSize),
		workers:   workers,
		generator: generator,
	}
}

func (s *Service) Run() {
	for i := 0; i < int(s.workers); i++ {
		s.runSingleWorker()
	}
}

func (s *Service) AddTask(in CreateTaskInput) (string, error) {
	jobID := s.generator.UniqueID()
	in.ID = jobID
	select {
	case s.jobBuffer <- in:
		s.storage.Set(in.ID, JobResult{
			Status: statusNew,
		})
	default:
		return "", errorx.ErrTooManyJobs
	}

	return jobID, nil
}

func (s *Service) GetTask(id string) (JobResult, error) {
	data, found := s.storage.GetByKey(id)
	if !found {
		return JobResult{}, errorx.ErrJobNotFound
	}
	res, ok := data.(JobResult)
	if !ok {
		s.logger.Fatal("wrong job structure")
	}
	return res, nil
}
