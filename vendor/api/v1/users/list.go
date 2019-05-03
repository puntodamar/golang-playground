package users

import (
	"github.com/gin-gonic/gin"
	. "configs/database"
	model_formatter "helpers/formatters/model"
	"strconv"

	//"strconv"
	//."helpers"
	. "models/v1"
	. "helpers"
)

func List(c *gin.Context) {
	db 			:= Db
	page, _ 	:= strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ 	:= strconv.Atoi(c.DefaultQuery("limit", "30"))

	formatter 	:= &[]model_formatter.User{}
	result 		:= Paginate(Param{
		DB 		: db.Scopes(ListUsers),
		Page 	: page,
		Limit 	: limit,
		Model	: &formatter,
	})


	//db.Scopes(ListUsers).Find(&result)

	c.JSON(200, gin.H{"data" : result})
}