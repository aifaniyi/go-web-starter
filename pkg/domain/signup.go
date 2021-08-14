package domain

import (
	"fmt"

	"github.com/aifaniyi/sample/pkg/entity"
	"github.com/aifaniyi/sample/pkg/exception"
	uuid "github.com/satori/go.uuid"
)

func (s *Domain) Signup(email, password string) error {
	var user *entity.User
	var err error

	if user, err = s.repo.GetUserRepo().ReadByEmail(email); err != nil {
		return &exception.ServerError{Err: err}
	}

	if user != nil {
		return &exception.ClientError{Err: fmt.Errorf("user with email %s already exists", email)}
	}

	if _, err = s.repo.GetUserRepo().Create(&entity.User{
		ID:       uuid.NewV4(),
		Email:    email,
		Password: password,
	}); err != nil {
		return err
	}

	return nil
}
