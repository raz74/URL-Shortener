package repository

import (
	"gorm.io/gorm"
	"shortened_link/model"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetByEmail(Email string) error
}

type UserRepositoryImpl struct {
	PostgresDb *gorm.DB
}

func (u *UserRepositoryImpl) CreateUser(user *model.User) error {
	return nil
}

func (u *UserRepositoryImpl) GetByEmail(Email string) error {
	return nil
}
