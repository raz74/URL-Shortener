package repository

import "shortened_link/model"

type TokenRepo interface {
	Create(user *model.User) (*model.SessionCookie, error)
	Get(token string) (*model.SessionCookie, error)
}
