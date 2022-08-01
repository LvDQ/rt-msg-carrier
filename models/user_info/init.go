package userinfo

import (
	"rt-msg-carrier/models/gen/dal/model"
	"rt-msg-carrier/tools"
)

var db = tools.GetDB()

func init() {
	db.AutoMigrate(&User{}, &Favorite{}, &model.Customer{})
}
