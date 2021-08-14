package domain

import (
	"github.com/aifaniyi/sample/pkg/repository"
)

type Domain struct {
	repo repository.Service
}

func NewDomain(repo repository.Service) *Domain {
	return &Domain{repo}
}
