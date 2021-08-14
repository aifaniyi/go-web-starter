package user

import (
	"fmt"

	"github.com/aifaniyi/sample/pkg/entity"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type RepoImpl struct {
	conn *gorm.DB
}

func NewRepoImpl(conn *gorm.DB) *RepoImpl {
	return &RepoImpl{conn}
}

func (r *RepoImpl) Create(user *entity.User) (*entity.User, error) {
	if res := r.conn.Create(user); res.Error != nil {
		return nil, fmt.Errorf("error creating user: %v", res.Error)
	}
	return user, nil
}

func (r *RepoImpl) ReadByEmail(email string) (*entity.User, error) {
	var user entity.User
	res := r.conn.Where("email = ?", email).Find(&user)
	if res.Error != nil {
		return nil, fmt.Errorf("error reading user by email: %v", res.Error)
	}
	if res.RowsAffected < 1 {
		return nil, nil
	}
	return &user, nil
}

func (r *RepoImpl) Update(user *entity.User) (*entity.User, error) {
	if res := r.conn.Model(&user).Updates(user); res.Error != nil {
		return nil, fmt.Errorf("error updating user: %v", res.Error)
	}
	return user, nil
}

func (r *RepoImpl) Delete(id uuid.UUID) error {
	if res := r.conn.Delete(&entity.User{ID: id}); res.Error != nil {
		return fmt.Errorf("error deleting user: %v", res.Error)
	}
	return nil
}
