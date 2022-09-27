package service

import (
	"startGo/internal/data"
	"startGo/internal/repository"
)

type UserService interface {
	Create(data *data.User) (int, error)
}

type userService struct {
	repository *repository.UserPg
}

func createUserService(r *repository.UserPg) *userService {
	return &userService{
		repository: r,
	}
}

func (user *userService) Create(data *data.User) (int, error) {

	(*user).repository.Create()

	return 1, nil
}
