package service

import (
	todo "github.com/prosto-kosmos/RestAPI_Go"
	"github.com/prosto-kosmos/RestAPI_Go/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(tocken string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	// GetAll(userId int) ([]todo.TodoList, error)
	// GetById(userId, listId int) (todo.TodoList, error)
	// Delete(userId, listId int) error
	// Update(userId, listId int, input todo.UpdateListInput) error
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
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
