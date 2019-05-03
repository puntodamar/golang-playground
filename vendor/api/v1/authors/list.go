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

	formatter 	:= &[]model.Author{}
	authors 	:= Paginate(Param{
		DB 		:
			db.
			Scopes(model.Author{}.ListOfAuthorsComplete).
			Preload("Books.Tags"),
		Page 	: page,
		Limit 	: limit,
		Model	: &formatter,
	})

	c.JSON(200, gin.H{"data" : &authors})
}
