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

func GetDetailUserNative(db *gorm.DB, id string) model.UserDetail {
	var usersDetail model.UserDetail
	db.Raw("select a.username as username, a.name as name, b.status as statusemp from t_user a, t_status b "+
		"where a.userid = ? and a.statusid = b.statusid", id).Scan(&usersDetail)

	return usersDetail
}
