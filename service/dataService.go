package service

import (
	"fmt"
	"gogin-practice/constant"
	"gogin-practice/model"
	io "gogin-practice/model/io"
	"gogin-practice/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Service GetUsers called ")

		users := repository.GetUserAll(db)

		var response io.Response
		for _, user := range users {
			var detail io.Detail
			detail.Id = user.Userid
			detail.Info = user.Username
			response.Data = append(response.Data, detail)
		}

		response.Status = constant.TRUE
		response.Message = "Success"

		c.JSON(http.StatusOK, response)
	}
}

func GetDetailUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := repository.GetDetailUser(db, c.Param("id"))

		var response io.Response

		var detail io.Detail
		detail.Id = user.Userid
		detail.Info = user.Username
		response.Data = append(response.Data, detail)

		response.Status = constant.TRUE
		response.Message = "Success"

		c.JSON(http.StatusOK, response)
	}
}

func SearchDetailUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := repository.GetDetailUser(db, c.Query("id"))

		var response io.Response

		var detail io.Detail
		detail.Id = user.Userid
		detail.Info = user.Username
		response.Data = append(response.Data, detail)

		response.Status = constant.TRUE
		response.Message = "Success"

		c.JSON(http.StatusOK, response)
	}
}

func LoginForm() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(200, gin.H{
			"status":   "posted",
			"username": username,
			"password": password,
		})

	}
}

type Userdata struct {
	Id       string `json: id binding:"required`
	Username string `json: username binding:"required`
	Name     string `json: name binding:"required`
}

func SetUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Service SetUser called")
		var userData Userdata
		c.BindJSON(&userData)

		user := model.TUser{Userid: userData.Id, Username: userData.Username, Name: userData.Name}
		result := db.Create(&user)

		var response io.Response
		if result.RowsAffected >= 1 {
			response.Status = constant.TRUE
			response.Message = "Success"

		} else {
			response.Status = constant.FALSE
			response.Message = "Failed"
		}

		c.JSON(http.StatusOK, response)
	}
}
