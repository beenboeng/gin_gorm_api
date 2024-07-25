package pkg

import (
	"fmt"
	"go_api/db"
	"go_api/internal/schemas"
	"strings"
)

type ClaimsObjectData struct {
	ID       int
	UserName string
	Session  string
	Exp      int
}

func ValidateUserSession(userReq ClaimsObjectData) bool {

	var users = schemas.User{}
	res := db.DB.Where(&schemas.User{UserName: userReq.UserName}).First(&users)

	if res.Error != nil {
		fmt.Println("Error in ValidateUserSession")
		fmt.Println(res.Error)
		return false
	}

	return strings.Contains(users.LoginSession, userReq.Session)
}
