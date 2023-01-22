package userRepository

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"shortened_link/model"
)

type UserRepositoryImpl struct {
	PostgresDb *gorm.DB
}

func (u *UserRepositoryImpl) CreateUser(user *model.User) error {
	return u.PostgresDb.Create(&user).Error
}

func (u *UserRepositoryImpl) CheckUniqueEmail(Email string) error {
	var user model.User
	err := u.PostgresDb.Where("email=?", Email).Find(&user).RowsAffected
	if err > 0 {
		return echo.ErrForbidden
	}
	return nil
}

func (u *UserRepositoryImpl) GetUserByEmail(Email string) (*model.User, error) {
	var user model.User
	err := u.PostgresDb.Where("email=?", Email).Find(&user).Error
	return &user, err
}
