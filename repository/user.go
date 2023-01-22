package repository

import (
	"shortened_link/model"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	CheckUniqueEmail(Email string) error
	GetUserByEmail(Email string) (*model.User, error)
}
