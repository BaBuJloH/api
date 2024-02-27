package service

import (
	"api"
	"api/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const salt = "1f1hno9q8fqil3n2g29gvbmas"

type AuthService struct {
	repo repository.Autorization
}

func NewAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user api.User) (int, error) {
	user.Password = genegatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func genegatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
