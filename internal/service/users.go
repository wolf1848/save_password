package service

import (
	"startGo/internal/data"
	"startGo/internal/repository"
)

type UserService interface {
	Create(data *data.User) (*data.User, error)
}

type userService struct {
	repository *repository.Repository
}

func createUserService(r *repository.Repository) *userService {
	return &userService{
		repository: r,
	}
}

func (user *userService) Create(data *data.User) (*data.User, error) {

	if id, err := user.repository.UserPg.Create(data); err != nil {
		return nil, err
	} else {
		data.Id = id
		return data, nil
	}

}
