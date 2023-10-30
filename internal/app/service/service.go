package service

import (
	"github.com/WORUS/arithmetic-progression/internal/app/repository"
	"github.com/WORUS/arithmetic-progression/internal/app/task"
)

type Task interface {
	SetTask()
	GetTask()
}

type Service struct {
	repository *repository.Repository
	queue      chan task.TaskInput
	goroutines chan bool
}

func NewService(repo *repository.Repository, que chan task.TaskInput, gorouts chan bool) *Service {
	return &Service{
		repository: repo,
		queue:      que,
		goroutines: gorouts,
	}
}
