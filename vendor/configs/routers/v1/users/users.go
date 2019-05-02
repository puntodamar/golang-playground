package users

import (
	"api/v1/users"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	u := route.Group("/api/v1/users")
	{
		u.GET("/", 		users.List)
	}
}
