package main

import (
	"go_api/db"
	"go_api/internal/routers"
	"go_api/pkg"
	"os"
)

func init() {
	pkg.LoadEnvFile()
	db.DatabaseConnection()
}

func main() {
	var port = os.Getenv("PORT")
	var base_id = os.Getenv("API_BASE_URL")
	router := routers.SetupRouter()
	router.Run(base_id + port)
}
