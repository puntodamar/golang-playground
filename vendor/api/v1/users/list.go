package users

import (
	"github.com/gin-gonic/gin"
	model_formatter "helpers/formatters/model/user"

	"strconv"

	. "configs/database"
	. "helpers"
	model "models/v1"
)

func List(c *gin.Context) {
	db 			:= Db
	page, _ 	:= strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ 	:= strconv.Atoi(c.DefaultQuery("limit", "30"))


	users 		:= Paginate(Param{
		DB 		: db.Scopes(model.ListUsers),
		Page 	: page,
		Limit 	: limit,
		Model	: model_formatter.User{},
	})

	c.JSON(200, gin.H{"data" : &users})
}