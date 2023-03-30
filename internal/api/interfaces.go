package api

import (
	"project/internal/services/task"
)

type taskService interface {
	AddTask(in task.CreateTaskInput) (string, error)
	GetTask(id string) (task.JobResult, error)
	Run()
}
