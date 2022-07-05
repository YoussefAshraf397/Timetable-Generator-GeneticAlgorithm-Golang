package Seeders

import (
	"github.com/brianvoe/gofakeit/v6"
	"go-starter/Models"
	"gorm.io/gorm"
)

func seedUsers(DB *gorm.DB) {
	DB.Create(&Models.User{
		Username: "Youssef Ashraf",
		Email:    "youssef@youssef.com",
		Password: "12345678",
		Group:    "admin",
		Token:    "adminyoussef",
	})
	DB.Create(&Models.User{
		Username: gofakeit.Name(),
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, false, false, 8),
		Group:    "user",
		Token:    "useryoussef",
	})
}
