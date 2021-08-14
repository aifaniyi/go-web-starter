package user

import (
	"github.com/aifaniyi/sample/pkg/entity"
	uuid "github.com/satori/go.uuid"
)

type Repo interface {
	Create(user *entity.User) (*entity.User, error)
	ReadByEmail(email string) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(id uuid.UUID) error
}
