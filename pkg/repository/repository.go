package repository

import (
	"github.com/jmoiron/sqlx"
	todo "github.com/prosto-kosmos/RestAPI_Go"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// NewRepository - конструктор структуры Repository
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
