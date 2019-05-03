package users

import (
	mid "api/middlewares/v1"
	"api/v1/users"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {

	u := route.Group("/api/v1/users", mid.JWT())
	{
		u.GET("/", users.List)
	}
}
