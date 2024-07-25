package routers

import (
	usersCon "go_api/internal/controllers/user"
	"go_api/internal/middleware"
	usersServ "go_api/internal/services/users"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	usersService := usersServ.NewUserService()
	usersController := usersCon.NewUserController(usersService)

	public := router.Group("/api")
	{
		public.POST("/users/register", usersController.Register)
		public.POST("/users/login", usersController.Login)
	}

	protected := router.Group("/api")
	protected.Use(middleware.Middleware())
	{
		protected.GET("/users", usersController.GetAllUsers)
		protected.POST("/users/add", usersController.CreateUser)
	}

	return router
}
