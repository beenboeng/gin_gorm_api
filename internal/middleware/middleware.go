package middleware

import (
	"fmt"
	"go_api/pkg"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if isCheckReqAuthentication(c) {
			c.Next()
		} else {
			abortRequest(c)
		}
	}
}

func abortRequest(c *gin.Context) {
	var response = pkg.ResBuilder(http.StatusUnauthorized, "unauthorized", pkg.Null())
	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func isCheckReqAuthentication(c *gin.Context) bool {

	fullToken := c.GetHeader("Authorization")
	if strings.TrimSpace(fullToken) != "" {
		trimedToken := strings.TrimSpace(fullToken)
		if len(trimedToken) > 10 {
			userToken := strings.Split(trimedToken, " ")[1]

			tokenValidate, err := validateToken(c, userToken)

			if err != nil {
				fmt.Println("Error in validate")
				return false
			}
			return tokenValidate

		} else {
			return false
		}

	} else {
		fmt.Println("Authorization is empty")
		return false
	}
}

type ClaimsObject struct {
	Exp          int
	ID           int    `json:"id"`
	UserName     string `json:"username"`
	LoginSession string `json:"login_session"`
	jwt.RegisteredClaims
}

func validateToken(c *gin.Context, tokenStr string) (val bool, err error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET_KEY"))

	token, err := jwt.ParseWithClaims(tokenStr, &ClaimsObject{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return false, err
	} else if claimsData, ok := token.Claims.(*ClaimsObject); ok {

		var newObj = pkg.ClaimsObjectData{
			ID:       claimsData.ID,
			UserName: claimsData.UserName,
			Session:  claimsData.LoginSession,
			Exp:      claimsData.Exp,
		}

		//Set new context every time when request
		c.Set("userCtx", newObj)

		validateSession := pkg.ValidateUserSession(newObj)
		if !validateSession {
			return false, err
		}
	} else {
		fmt.Println("unknown claims type, cannot proceed")
		return false, err
	}
	return true, err
}
