package entity

import uuid "github.com/satori/go.uuid"

type User struct {
	ID       uuid.UUID `gorm:"primary_key"`
	Email    string
	Password string
}
