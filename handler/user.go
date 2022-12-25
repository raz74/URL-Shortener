package handler

import (
	"github.com/labstack/echo"
	"shortened_link/repository"
)

type UserHandler struct {
	r repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{
		r: repo,
	}
}

func (u *UserHandler) Login(c echo.Context) error {
	return nil
}

func (u *UserHandler) SignUp(c echo.Context) error {
	return nil
}
