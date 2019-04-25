package users

import (

	"github.com/gin-gonic/gin"

	"strconv"

	. "helpers"
	. "models/v1"
	. "configs/database"
)


func List(c *gin.Context) {
	db 			:= Db
	page, _ 	:= strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ 	:= strconv.Atoi(c.DefaultQuery("limit", "30"))

	var users []User

	result := Paginate(Param{
		DB 		: db
		Page 	: page,
		Limit 	: limit,
		Model	: &users,
	})

	c.JSON(200, result)

	db.Close()
}
