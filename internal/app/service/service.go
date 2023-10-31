package service

import (
	"github.com/WORUS/arithmetic-progression/internal/app/cache"
	"github.com/WORUS/arithmetic-progression/internal/app/task"
)

type Task interface {
	SetTask()
	GetTask()
}

type Service struct {
	cache      *cache.Cache
	queue      chan *task.Task
	goroutines chan bool
}

func NewService(cache *cache.Cache, que chan *task.Task, gorouts chan bool) *Service {
	return &Service{
		cache:      cache,
		queue:      que,
		goroutines: gorouts,
	}
}
