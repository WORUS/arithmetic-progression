package service

import "github.com/WORUS/arithmetic-progression/internal/app/repository"

type Task interface {
	SetTask()
	GetTask()
}

type Service struct {
	repo *repository.Repository
}
