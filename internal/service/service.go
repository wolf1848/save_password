package service

import "startGo/internal/repository"

type UserService interface {
	create() (int, error)
}

type Service struct {
	UserService
}

func CreateService(r *repository.Repository) *Service {
	return &Service{
		UserService: createUserService(&r.UserPg),
	}
}
