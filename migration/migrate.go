package main

import (
	"fmt"
	"go_api/db"
	"go_api/internal/schemas"
	"go_api/pkg"
)

func init() {
	pkg.LoadEnvFile()
	db.DatabaseConnection()
}

func main() {
	fmt.Println("Migrating ....")
	defer fmt.Println("Migrating completed")
	err := db.DB.AutoMigrate(&schemas.User{})

	if err != nil {
		fmt.Println("Error in migrate!")
		return
	}

}
