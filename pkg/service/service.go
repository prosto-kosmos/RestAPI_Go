package service

import "github.com/prosto-kosmos/RestAPI_Go/pkg/repository"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

// NewService - конструктор структуры Service
func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
