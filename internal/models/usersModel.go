package models

type UserInfo struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
}

type UserInfoRequest struct {
	FirstName string  `binding:"required"`
	LastName  string  `binding:"required"`
	UserName  string  `binding:"required"`
	Email     *string // A pointer to a string, allowing for null values
}

type UserInfoRegisterRequest struct {
	FirstName string  `binding:"required"`
	LastName  string  `binding:"required"`
	UserName  string  `binding:"required"`
	Email     *string // A pointer to a string, allowing for null values
	Password  string  `binding:"required"`
}

type UserLoginRequest struct {
	UserName string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}

type UserInfoLoginRespone struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserName     string `json:"username"`
	Email        string `json:"email"`
	LoginSession string `json:"login_session"`
	Token        string `json:"token"`
}

type UserContext struct {
	ID       int
	UserName string
	Session  string
	Exp      int
}
