package user

import (
	"context"

	"github.com/aifaniyi/sample/pkg/entity"
	uuid "github.com/satori/go.uuid"
)

type Repo interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	ReadByEmail(ctx context.Context, email string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
