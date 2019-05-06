package authors

import (
	"github.com/gin-gonic/gin"
	"strconv"

	. "helpers"
	. "configs/database"
	model "models/v1"
	//model_formatter "helpers/formatters/model"
)

func List(c *gin.Context) {
	db 			:= Db

	page, _ 	:= strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ 	:= strconv.Atoi(c.DefaultQuery("limit", "30"))

	authors 	:= Paginate(Param{
		DB 		:
			db.
			Scopes(model.CompleteAuthorList).
			Preload("Books.Tags"),
		Page 	: page,
		Limit 	: limit,
		Model	: &[]model.Author{},
	})

	c.JSON(200, gin.H{"data" : &authors})
}