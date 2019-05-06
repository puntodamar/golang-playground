package seeders

import (
	"api/v1/seeders"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {

	u := route.Group("/api/v1/seed")
	{
		u.GET("/", seeders.Seed)
	}
}
