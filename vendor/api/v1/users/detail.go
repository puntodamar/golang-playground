package users

import (
	. "configs/database"
	"github.com/gin-gonic/gin"
	model_formatter "helpers/formatters/model/user"
	"net/http"
)

func Detail(c *gin.Context) {
	db := Db

	user 	:= &model_formatter.User{}
	id 		:= c.Param("id")

	if err := db.Find(&user,id).Error; err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"message" : "can't find user"},
		)
	}

	c.JSON(http.StatusOK, user)
}
