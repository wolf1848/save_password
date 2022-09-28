package service

import (
	"errors"
	"os"
	"startGo/internal/data"
	"startGo/internal/repository"

	"golang.org/x/crypto/bcrypt"
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

	if err := user.hashGen(&data.Pwd); err != nil {
		return nil, err
	}

	if id, err := user.repository.UserPg.Create(data); err != nil {
		return nil, err
	} else {
		data.Id = id
		data.Pwd = ""
		return data, nil
	}

}

func (user *userService) hashGen(pwd *string) error {
	if len(*pwd) >= 8 {
		saltedBytes := []byte(os.Getenv("SERVER_SALT"))
		pwdBytes := []byte(*pwd)
		bytes := append(pwdBytes[:len(pwdBytes)/2], saltedBytes...)
		bytes = append(bytes, pwdBytes[len(pwdBytes)/2:]...)

		hashedBytes, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		*pwd = string(hashedBytes[:])
		return nil
	} else {
		return errors.New("длина поля \"Пароль\" меньше 8 символов")
	}

}
