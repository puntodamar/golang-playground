package users

import (
	"github.com/gin-gonic/gin"

	//"log"
	//"net/http"
	"fmt"

	. "configs/database"
	. "models/v1/user"
)

//type User struct {
//	ID        int   `json:"id"`
//	Name 	string `json:"name"`
//
//}

func List(c *gin.Context) {
	db 		:= Db
	var users []User

	if err := db.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}

	//if err := db.Find(&users).Error; err != nil {
	//	c.AbortWithStatus(http.StatusInternalServerError)
	//	fmt.Println(err)
	//} else {
	//	c.JSON(http.StatusOK, users)
	//	log.Println(users)
	//}

	db.Close()
}
