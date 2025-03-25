package repository

import (
	"database/sql"
	"gogin-practice/model"

	"gorm.io/gorm"
)

func GetUserAll(db *gorm.DB) []model.TUser {
	var users []model.TUser
	db.Find(&users)

	// fmt.Println("userid " + user.Userid)
	// fmt.Println("username " + user.Username)

	return users
}

func GetDetailUser(db *gorm.DB, id string) model.TUser {
	var users model.TUser
	db.Where("userid = @userid", sql.Named("userid", id)).Find(&users)

	return users
}
