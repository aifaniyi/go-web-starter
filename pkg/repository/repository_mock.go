package repository

import "github.com/aifaniyi/sample/pkg/repository/user"

type Mock struct {
	userRepo user.Repo
}

func NewServiceMock() *Mock {
	return &Mock{}
}

func (m *Mock) GetUserRepo() user.Repo {
	if m.userRepo == nil {
		m.userRepo = &user.RepoMock{}
	}
	return m.userRepo
}
