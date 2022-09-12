package service

import "startGo/internal/repository"

type userService struct {
	repository *repository.UserPg
}

func createUserService(r *repository.UserPg) *userService {
	return &userService{
		repository: r,
	}
}

func (user *userService) create() (int, error) {
	return 1, nil
}
