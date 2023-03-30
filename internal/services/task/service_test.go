package task

import (
	"project/internal/common/errorx"
	"project/internal/mock"
	"project/internal/services/task/storage/memomry"
	"project/pkg/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTask(t *testing.T) {
	log := logger.NewTestLogger()
	t.Run("successfully add task", func(t *testing.T) {
		fakeGenerator := mock.NewGenerator()
		id := "random-id"
		fakeGenerator.On("UniqueID").Return(id)
		storage := memomry.NewMemory()
		svc := NewService(log, storage, nil, fakeGenerator, 10, 1024)
		res, err := svc.AddTask(CreateTaskInput{})
		assert.NoError(t, err)
		assert.Equal(t, id, res)
		job, err := svc.GetTask(id)
		assert.NoError(t, err)
		assert.EqualValues(t, statusNew, job.Status)
	})

	t.Run("failed because of full buffer", func(t *testing.T) {
		fakeGenerator := mock.NewGenerator()
		id := "random-id"
		fakeGenerator.On("UniqueID").Return(id)
		storage := memomry.NewMemory()
		svc := NewService(log, storage, nil, fakeGenerator, 1, 1)
		_, err := svc.AddTask(CreateTaskInput{})
		assert.NoError(t, err)
		res, err := svc.AddTask(CreateTaskInput{})
		assert.ErrorIs(t, err, errorx.ErrTooManyJobs)
		assert.Empty(t, res)
	})
}

func TestGetTask(t *testing.T) {
	log := logger.NewTestLogger()
	t.Run("successfully get task", func(t *testing.T) {
		fakeGenerator := mock.NewGenerator()
		id := "random-id"
		fakeGenerator.On("UniqueID").Return(id)
		storage := memomry.NewMemory()
		svc := NewService(log, storage, nil, fakeGenerator, 10, 1024)
		_, err := svc.AddTask(CreateTaskInput{})
		assert.NoError(t, err)
		res, err := svc.GetTask(id)
		assert.NoError(t, err)
		assert.EqualValues(t, statusNew, res.Status)
	})

	t.Run("task not exists", func(t *testing.T) {
		fakeGenerator := mock.NewGenerator()
		id := "random-id"
		fakeGenerator.On("UniqueID").Return(id)
		storage := memomry.NewMemory()
		svc := NewService(log, storage, nil, fakeGenerator, 10, 1024)
		_, err := svc.GetTask(id)
		assert.ErrorIs(t, err, errorx.ErrJobNotFound)
	})
}
