package user

import (
	"context"

	"github.com/aifaniyi/sample/pkg/entity"
	uuid "github.com/satori/go.uuid"
)

type RepoMock struct {
	data []entity.User
}

func (r *RepoMock) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	user.ID = uuid.NewV4()
	r.data = append(r.data, *user)
	return user, nil
}

func (r *RepoMock) ReadByEmail(ctx context.Context, email string) (*entity.User, error) {
	for _, d := range r.data {
		if email == d.Email {
			return &d, nil
		}
	}

	return nil, nil
}

func (r *RepoMock) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	for index := range r.data {
		if r.data[index].ID.String() == user.ID.String() {
			r.data[index] = *user
			return user, nil
		}
	}

	return nil, nil
}

func (r *RepoMock) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
