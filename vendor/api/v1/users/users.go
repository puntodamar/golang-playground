package users

import (
	"github.com/gin-gonic/gin"
	"strconv"
	."helpers"
	. 				"models/v1"
	. 				"configs/database"
	model_formatter "helpers/formatters/model"
)


func List(c *gin.Context) {
	db 			:= Db
	page, _ 	:= strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ 	:= strconv.Atoi(c.DefaultQuery("limit", "30"))

	var users []User

	result := Paginate(Param{
		DB 		: db,
		Page 	: page,
		Limit 	: limit,
		Model	: &users,
	})

	val := db.Scopes(ListUsers).Find(&model_formatter.User{}).Value

	c.JSON(200, val)
}

func Login(c *gin.Context) {


}
