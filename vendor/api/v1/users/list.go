package users

import (
	"github.com/gin-gonic/gin"


	"strconv"

	. "models/v1"
	. "helpers"
	. "configs/database"
	model_formatter "helpers/formatters/model"
)

func List(c *gin.Context) {
	db 			:= Db
	page, _ 	:= strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ 	:= strconv.Atoi(c.DefaultQuery("limit", "30"))

	formatter 	:= &[]model_formatter.User{}
	users 		:= Paginate(Param{
		DB 		: db.Scopes(ListUsers),
		Page 	: page,
		Limit 	: limit,
		Model	: &formatter,
	})

	c.JSON(200, gin.H{"data" : &users})
}