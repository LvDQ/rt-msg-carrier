package tools

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}

func TestConn(t *testing.T) {
	db := GetDB()
	sqlDB, _ := db.DB()
	fmt.Println(sqlDB.Ping())

	var users []User
	db.First(&users)
	fmt.Printf("test main\n")
	t.Log("complete")
}
