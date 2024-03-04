package service

import (
	"api"
	"api/pkg/repository"
)

type Autorization interface {
	CreateUser(user api.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list api.TodoList) (int, error)
	GetAll(userId int) ([]api.TodoList, error)
	GetById(userId, listId int) (api.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input api.UpdateListInput) error
}

type TodoItem interface {
}

type Service struct {
	Autorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
		TodoList:     NewTodoListService(repos.TodoList),
	}
}
