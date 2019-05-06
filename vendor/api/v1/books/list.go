package books

import (
	"github.com/gin-gonic/gin"

	"strconv"

	. "configs/database"
	. "helpers"
	model "models/v1"
)

func List(c *gin.Context) {
	db 			:= Db
	page, _ 	:= strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ 	:= strconv.Atoi(c.DefaultQuery("limit", "30"))

	books 		:= Paginate(Param{
		DB 		: db.Scopes(model.BookList),
		Page 	: page,
		Limit 	: limit,
		Model	: &model.Book{},
	})

	c.JSON(200, gin.H{"data" : &books})
}