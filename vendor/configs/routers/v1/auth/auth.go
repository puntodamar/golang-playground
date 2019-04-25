package auth

import (
	"api/v1/authentication"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine){
	a := route.Group("/api/v1/auth")
	{
		a.POST("/login", authentication.Login)
	}
}
