package user

import (
	"github.com/aifaniyi/sample/pkg/entity"
	uuid "github.com/satori/go.uuid"
)

type RepoMock struct {
	data []entity.User
}

func (r *RepoMock) Create(user *entity.User) (*entity.User, error) {
	user.ID = uuid.NewV4()
	r.data = append(r.data, *user)
	return user, nil
}

func (r *RepoMock) ReadByEmail(email string) (*entity.User, error) {
	for _, d := range r.data {
		if email == d.Email {
			return &d, nil
		}
	}

	return nil, nil
}

func (r *RepoMock) Update(user *entity.User) (*entity.User, error) {
	for index := range r.data {
		if r.data[index].ID.String() == user.ID.String() {
			r.data[index] = *user
			return user, nil
		}
	}

	return nil, nil
}

func (r *RepoMock) Delete(id uuid.UUID) error {
	return nil
}
