package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/aifaniyi/env"
	"github.com/aifaniyi/sample/pkg/entity"
	"github.com/aifaniyi/sample/pkg/repository/user"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	defaultQueryTimeout = 2
)

type Impl struct {
	conn         *gorm.DB
	queryTimeout time.Duration
	userRepo     user.Repo
}

func (i *Impl) GetUserRepo() user.Repo {
	if i.userRepo == nil {
		i.userRepo = user.NewRepoImpl(i.conn, i.queryTimeout)
	}
	return i.userRepo
}

func NewServiceImpl() (*Impl, error) {
	host := env.LoadString("DB_HOST", "localhost")
	port := env.LoadString("DB_PORT", "5432")
	database := env.LoadString("DB_NAME", "sample")
	user := env.LoadString("DB_USER", "postgres")
	password := env.LoadString("DB_PASSWORD", "postgres")
	ssl := env.LoadBool("DB_SSL", false)
	queryTimeout := time.Duration(env.LoadInt("DB_QUERY_TIMEOUT", defaultQueryTimeout)) * time.Second

	conn, err := connect(host, port, database, user, password, ssl)
	if err != nil {
		return nil, err
	}

	gconn, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = migrate(gconn)
	if err != nil {
		return nil, err
	}

	return &Impl{
		conn:         gconn,
		queryTimeout: queryTimeout,
	}, nil
}

func migrate(conn *gorm.DB) error {
	entities := []interface{}{
		&entity.User{},
	}

	for _, model := range entities {
		if !conn.Migrator().HasTable(model) {
			if err := conn.AutoMigrate(model); err != nil {
				return err
			}
		}
	}
	return nil
}

// connect : create a database connection
func connect(host, port, database, user, password string, ssl bool) (*sql.DB, error) {
	var dbInfo string
	if ssl {
		dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=enable",
			host, port, user, password, database)
	} else {
		dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, database)
	}

	return sql.Open("postgres", dbInfo)
}
