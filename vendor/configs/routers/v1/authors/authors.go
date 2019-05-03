package authors

import (
	"api/v1/authors"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine){
	a := route.Group("/api/v1/authors")
	{
		a.GET("/", authors.List)
	}
}
