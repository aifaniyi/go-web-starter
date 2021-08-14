package mock

import "github.com/aifaniyi/sample/pkg/repository"

func Repository() repository.Service {
	return repository.NewServiceMock()
}
