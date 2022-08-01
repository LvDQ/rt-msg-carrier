package userinfo

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
}

func GetUserDetailById(id uint) (user User) {
	db.Where("id = ?", id).First(&user)
	return user
}

func CreateNewUser() (user User) {
	user = User{Name: "ky", Email: "kyky@gmail.com", Gender: "male"}
	result := db.Create(&user)
	fmt.Println(user.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
	return user
}
