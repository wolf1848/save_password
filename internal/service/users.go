package service

import (
	"errors"
	"os"
	"regexp"
	"startGo/internal/data"
	"startGo/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(data *data.User) (*data.User, error)
	LogIn(data *data.User) (*data.User, error)
}

type userService struct {
	repository *repository.Repository
}

func salted(pwd *string) (bytes []byte) {
	saltedBytes := []byte(os.Getenv("SERVER_SALT"))
	pwdBytes := []byte(*pwd)
	bytes = append(pwdBytes[:len(pwdBytes)/2], saltedBytes...)
	bytes = append(bytes, pwdBytes[len(pwdBytes)/2:]...)
	return
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

func (user *userService) LogIn(data *data.User) (*data.User, error) {
	id, pwd, full_name, err := user.repository.UserPg.CheckUser(data)
	if err != nil {
		return nil, err
	}

	incoming := salted(&data.Pwd)
	existing := []byte(pwd)
	if err := bcrypt.CompareHashAndPassword(existing, incoming); err != nil {
		return nil, err
	}

	data.Id = id
	data.Pwd = ""
	data.FullName = full_name

	return data, nil
}

func (user *userService) hashGen(pwd *string) error {

	secure := true
	tests := []string{".{8,}", "[a-z]", "[A-Z]", "[0-9]", "[^\\d\\w]"} // Длина 8 символов, буквы нижний верхний регистр, символы.
	for _, test := range tests {
		t, _ := regexp.MatchString(test, *pwd)
		if !t {
			secure = false
			break
		}
	}

	if secure {

		bytes := salted(pwd)

		hashedBytes, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		*pwd = string(hashedBytes[:])
		return nil
	} else {
		return errors.New("поле \"Пароль\" не соответствует требованиям: длина не менее 8 символов, цифры и буквы верхнего и нижнего регистра, спецсимволы")
	}

}
