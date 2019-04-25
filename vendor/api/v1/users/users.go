package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	. "configs/database"
	. "models/v1/user"
)

func List(c *gin.Context) {

	db := Db

	user := User{}

	if err := db.First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
		log.Println(user)
	}

	db.Close()
}
