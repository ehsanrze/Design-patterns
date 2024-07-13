package database

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      float64
	LastName  string
	Credit    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func UpdateUser(userId int, user User) {

}
