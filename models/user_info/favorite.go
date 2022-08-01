package userinfo

import (
	"fmt"
	"rt-msg-carrier/tools"

	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserID int
	Menu   string
	User   User
}

func GetFavoriteDetailsByUserid(id int) (favorate Favorite) {
	db := tools.GetDB()
	favorite := Favorite{}
	db.Preload("User").Find(&favorite)
	fmt.Println(favorite)
	return favorite
}
