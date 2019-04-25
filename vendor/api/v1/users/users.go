package users

import (
	"github.com/gin-gonic/gin"

	"fmt"

	. "configs/database"
	. "models/v1/user"
)


func List(c *gin.Context) {
	db 		:= Db
	var users []User

	if err := db.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}

	db.Close()
}
