package users

import (
	"fmt"
	"go_api/db"
	"go_api/internal/models"
	"go_api/internal/schemas"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (t UserService) FindAll() (data []schemas.User, err error) {

	var users []schemas.User
	res := db.DB.Find(&users)
	return users, res.Error
}

func (t UserService) CreateUser(userReq models.UserInfoRequest) (res models.UserInfo, err error) {

	var emptyEmail = ""
	user := schemas.User{
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		UserName:  userReq.UserName,
		Email:     &emptyEmail,
	}

	if userReq.Email != nil {
		user.Email = userReq.Email
	}

	result := db.DB.Create(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	resData := models.UserInfo{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     *user.Email,
	}

	return resData, result.Error
}

func (t UserService) Register(userReq models.UserInfoRegisterRequest) (res models.UserInfo, err error) {

	var emptyEmail = ""
	user := schemas.User{
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		UserName:  userReq.UserName,
		Password:  userReq.Password,
		Email:     &emptyEmail,
	}

	if userReq.Email != nil {
		user.Email = userReq.Email
	}

	result := db.DB.Create(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	resData := models.UserInfo{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     *user.Email,
	}

	return resData, result.Error
}

func (t UserService) Login(userReq models.UserLoginRequest) (data schemas.User, err error) {

	var users = schemas.User{}
	res := db.DB.Where(&schemas.User{UserName: userReq.UserName}).First(&users)

	if res.Error != nil {
		fmt.Println("In service login")
		fmt.Println(res.Error)
	}

	return users, res.Error
}

func (u UserService) SetUserNewSession(userName, userSession string) (val bool, err error) {

	var user = schemas.User{
		UserName: userName,
	}
	res := db.DB.Model(&user).Update("login_session", userSession)
	if res.Error != nil {
		return false, res.Error
	}

	return true, res.Error
}

func (u UserService) GenerateToken(userInfo schemas.User, session string) (userRespone models.UserInfoLoginRespone, err error) {

	jwtSecret := []byte(os.Getenv("JWT_SECRET_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":            userInfo.ID,
		"username":      userInfo.UserName,
		"login_session": session,
		"exp":           time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)

	var userRes = models.UserInfoLoginRespone{}
	if err != nil {
		fmt.Println(err.Error())
		return userRes, err
	}

	userRes = models.UserInfoLoginRespone{
		ID:           userInfo.ID,
		FirstName:    userInfo.FirstName,
		LastName:     userInfo.LastName,
		UserName:     userInfo.UserName,
		Email:        *userInfo.Email,
		LoginSession: session,
		Token:        tokenString,
	}

	return userRes, err

}
