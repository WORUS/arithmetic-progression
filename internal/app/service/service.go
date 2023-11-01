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
	qready     chan bool
	queue      []*task.Task
	goroutines chan bool
}

func NewService(cache *cache.Cache, qr chan bool, gorouts chan bool, que []*task.Task) *Service {
	return &Service{
		cache:      cache,
		qready:     qr,
		queue:      que,
		goroutines: gorouts,
	}
}
