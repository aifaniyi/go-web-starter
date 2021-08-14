package repository

import "github.com/aifaniyi/sample/pkg/repository/user"

type Service interface {
	GetUserRepo() user.Repo
}
