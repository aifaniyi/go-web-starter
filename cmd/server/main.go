package main

import (
	"log"

	"github.com/aifaniyi/sample/pkg/config"
	"github.com/aifaniyi/sample/pkg/repository"
)

func main() {
	conf := config.LoadConfig()
	run(conf)
}

func run(conf *config.Config) {
	repo := createRepo()

	svr := NewServer(repo)
	svr.start(conf.Port)
}

var createRepo = func() repository.Service {
	repo, err := repository.NewServiceImpl()
	if err != nil {
		log.Fatalf("error starting postgres server: %v", err)
	}

	return repo
}
