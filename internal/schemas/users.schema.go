package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           int `gorm:"primaryKey;autoIncrement:true"`
	FirstName    string
	LastName     string
	UserName     string  `gorm:"primaryKey;autoIncrement:true;unique:true"`
	Email        *string // A pointer to a string, allowing for null values
	Password     string
	LoginSession string
	StatusId     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
