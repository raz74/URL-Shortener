package handler

import (
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"shortened_link/handler/request"
	"shortened_link/model"
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

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(Password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(Password))
	return err
}

func (u *UserHandler) SignUp(c echo.Context) error {
	var req request.UserRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.ErrBadRequest
	}
	hash, _ := hashPassword(req.Password)

	User := model.User{
		Name:     req.Name,
		Password: hash,
		Email:    req.Email,
	}

	err = u.r.CheckUniqueEmail(req.Email)
	if err != nil {
		return c.JSON(http.StatusForbidden, "This email is already exist, try another!")
	}
	err = u.r.CreateUser(&User)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, "sign up successfully.")
}

func (u *UserHandler) Login(c echo.Context) error {
	var req request.UserRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.ErrBadRequest
	}
	user, err := u.r.GetUserByEmail(req.Email)
	if err != nil {
		return echo.ErrNotFound
	}
	check := checkPasswordHash(req.Password, user.Password)
	if check != nil {
		return c.JSON(http.StatusForbidden, "password is wrong! try again")
	}
	return c.JSON(http.StatusOK, user)
}
