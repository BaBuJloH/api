package repository

import (
	"api"

	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreateUser(api.User) (int, error)
	GetUser(username, password string) (api.User, error)
}

type TodoList interface {
	Create(userId int, list api.TodoList) (int, error)
}

type TodoItem interface {
}

type Repository struct {
	Autorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthPostgres(db),
		TodoList:     NewTodoListPostgres(db),
	}
}
