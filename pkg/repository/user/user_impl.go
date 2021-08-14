package user

import (
	"context"
	"fmt"
	"time"

	"github.com/aifaniyi/sample/pkg/entity"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type RepoImpl struct {
	conn         *gorm.DB
	queryTimeout time.Duration
}

func NewRepoImpl(conn *gorm.DB, queryTimeout time.Duration) *RepoImpl {
	return &RepoImpl{
		conn:         conn,
		queryTimeout: queryTimeout,
	}
}

func (r *RepoImpl) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	if res := r.conn.WithContext(ctx).Create(user); res.Error != nil {
		return nil, fmt.Errorf("error creating user: %v", res.Error)
	}
	return user, nil
}

func (r *RepoImpl) ReadByEmail(ctx context.Context, email string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	var user entity.User
	res := r.conn.WithContext(ctx).Where("email = ?", email).Find(&user)
	if res.Error != nil {
		return nil, fmt.Errorf("error reading user by email: %v", res.Error)
	}
	if res.RowsAffected < 1 {
		return nil, nil
	}
	return &user, nil
}

func (r *RepoImpl) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	if res := r.conn.WithContext(ctx).Model(&user).Updates(user); res.Error != nil {
		return nil, fmt.Errorf("error updating user: %v", res.Error)
	}
	return user, nil
}

func (r *RepoImpl) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	if res := r.conn.WithContext(ctx).Delete(&entity.User{ID: id}); res.Error != nil {
		return fmt.Errorf("error deleting user: %v", res.Error)
	}
	return nil
}
