package users

import (
	"api/middlewares"
	"api/v1/users"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {

	u := route.Group("/api/v1/users", middlewares.JWT())
	{
		u.GET("/", users.List)
		u.GET("/:id", users.Detail)
	}
}
