package users

import (
	"fmt"
	"go_api/internal/models"
	"go_api/internal/services/users"
	"go_api/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	UserService *users.UserService
}

func NewUserController(s *users.UserService) *UserController {
	return &UserController{
		UserService: s,
	}
}

func (u *UserController) GetAllUsers(c *gin.Context) {
	users, err := u.UserService.FindAll()

	//Example of using Context

	// userCtx, ctxerr := c.Get("userCtx")
	// if !ctxerr {
	// 	fmt.Println("Error get User Context")
	// }

	// var hehe = userCtx.(pkg.ClaimsObjectData)

	// fmt.Println(hehe.UserName, "userCtx")

	if err != nil {
		fmt.Println("Controller: Error in get all users!")
		var respone = pkg.ResBuilder(http.StatusBadRequest, "Error", pkg.Null())
		c.JSON(http.StatusBadRequest, respone)
	}

	responeData := []models.UserInfo{}
	for _, val := range users {
		responeData = append(responeData, models.UserInfo{
			ID:        val.ID,
			FirstName: val.FirstName,
			LastName:  val.LastName,
			UserName:  val.UserName,
			Email:     *val.Email,
		})
	}

	var successRespone = pkg.ResBuilder(http.StatusOK, "Success", responeData)
	c.JSON(http.StatusOK, successRespone)
}

func (u *UserController) CreateUser(c *gin.Context) {

	var body models.UserInfoRequest
	check_err := c.ShouldBindJSON(&body)
	if check_err != nil {
		var respone = pkg.ResBuilder(http.StatusBadRequest, check_err.Error(), pkg.Null())
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	users, err := u.UserService.CreateUser(body)
	if err != nil {
		fmt.Println("Controller: Error in create user!")
		var respone = pkg.ResBuilder(http.StatusBadRequest, "Error", pkg.Null())
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	var successRespone = pkg.ResBuilder(http.StatusOK, "Success", users)
	c.JSON(http.StatusOK, successRespone)
}

func (u *UserController) Register(c *gin.Context) {

	var body models.UserInfoRegisterRequest
	check_err := c.ShouldBindJSON(&body)
	if check_err != nil {
		var respone = pkg.ResBuilder(http.StatusBadRequest, check_err.Error(), pkg.Null())
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	users, err := u.UserService.Register(body)
	if err != nil {
		fmt.Println("Controller: Error in Register user!")
		var respone = pkg.ResBuilder(http.StatusBadRequest, "Register Error", pkg.Null())
		c.JSON(http.StatusBadRequest, respone)
		return
	}

	var successRespone = pkg.ResBuilder(http.StatusOK, "Success", users)
	c.JSON(http.StatusOK, successRespone)
}

func (u *UserController) Login(c *gin.Context) {

	var params models.UserLoginRequest
	check_err := c.ShouldBindJSON(&params)
	if check_err != nil {
		var respone = pkg.ResBuilder(http.StatusUnauthorized, check_err.Error(), pkg.Null())
		c.JSON(http.StatusUnauthorized, respone)
		return
	}

	// Check if user exist
	userInfo, err := u.UserService.Login(params)
	if err != nil {
		fmt.Println("Controller: Error in get users!")
		var respone = pkg.ResBuilder(http.StatusUnauthorized, "Login Fail", pkg.Null())
		c.JSON(http.StatusUnauthorized, respone)
		return
	}

	// Check user password
	if userInfo.Password != params.Password {
		var respone = pkg.ResBuilder(http.StatusUnauthorized, "Login Fail", pkg.Null())
		c.JSON(http.StatusUnauthorized, respone)
		return
	}

	// Set new session for user
	newUserSession := uuid.New()
	setSession, err := u.UserService.SetUserNewSession(userInfo.UserName, newUserSession.String())
	if err != nil || !setSession {
		var respone = pkg.ResBuilder(http.StatusUnauthorized, "Login Fail", pkg.Null())
		c.JSON(http.StatusUnauthorized, respone)
		return
	}

	// Generate user token
	userRespone, tokenErr := u.UserService.GenerateToken(userInfo, newUserSession.String())
	if tokenErr != nil {
		var respone = pkg.ResBuilder(http.StatusUnauthorized, "Login Fail", pkg.Null())
		c.JSON(http.StatusUnauthorized, respone)
		return
	}

	var successRespone = pkg.ResBuilder(http.StatusOK, "Success", userRespone)
	c.JSON(http.StatusOK, successRespone)
}
