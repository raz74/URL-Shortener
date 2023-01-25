package token

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"shortened_link/model"
	"time"
)

type PostTokenRepoImp struct {
	PostgresDb *gorm.DB
}

func (t *PostTokenRepoImp) Create(user *model.User) (*model.SessionCookie, error) {
	//create a new random cookie session
	cookieToken := uuid.NewString()
	expiresAt := time.Now().Add(7 * 24 * time.Second)
	var token model.SessionCookie
	token = model.SessionCookie{
		UserID: user.Id,
		Value:  cookieToken,
		Expire: expiresAt,
	}
	err := t.PostgresDb.Create(&token).Error

	return &token, err
}

func (t *PostTokenRepoImp) Get(header string) (*model.SessionCookie, error) {
	substr := header[6:]
	println(substr)
	var token model.SessionCookie
	err := t.PostgresDb.Where("value=?", substr).Find(&token).Error
	if err != nil {
		fmt.Println("this session is not exist")
		return nil, err
	}
	if token.Expire.Before(time.Now()) {
		t.PostgresDb.Where("value=?", substr).Delete(&token)
		return nil, err
	}
	return &token, err
}
