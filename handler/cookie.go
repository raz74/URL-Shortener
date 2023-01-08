package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"shortened_link/model"
	"time"
)

func GenerateSession(user *model.User) string {
	//create a new random cookie session
	cookieToken := uuid.NewString()
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	cookieMap[cookieToken] = model.SessionCookie{
		Expire: expiresAt,
		Name:   user.Name,
	}
	return cookieToken
}

func CheckHeaderAuthorize(header string) error {
	substr := header[6:]
	println(substr)
	userSession, exist := cookieMap[substr]
	if !exist {
		fmt.Println("this session is not exist")
		return echo.ErrUnauthorized

	}
	if userSession.IsExpire() {
		delete(cookieMap, substr)
		return echo.ErrUnauthorized
	}
	return nil
}
